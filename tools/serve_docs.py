# tools/serve_docs.py
#!/usr/bin/env python3
import logging
import sys
import threading
import webbrowser
from http.server import SimpleHTTPRequestHandler, HTTPServer
from pathlib import Path
import socket
from functools import partial

from tools.common import setup_logging, ensure_directory_exists

def find_free_port():
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind(("", 0))
        return s.getsockname()[1]

def find_docs_directory():
    logger = setup_logging()
    
    possible_dirs = [
        "docs/html",
        "docs/dart/html",
        "docs/api/html"
    ]
    
    for dir_path in possible_dirs:
        path = Path(dir_path)
        if path.exists() and (path / "index.html").exists():
            logger.info(f"Found documentation at {path}")
            return str(path)
    
    logger.warning("HTML documentation not found. Generating docs first...")
    from tools.generate_docs_artifacts import main as generate_docs
    generate_docs()
    
    html_dir = Path("docs/html")
    if html_dir.exists() and (html_dir / "index.html").exists():
        return str(html_dir)
    
    logger.error("Could not find or generate HTML documentation")
    sys.exit(1)

def serve_documentation(port: int = 8000):
    logger = setup_logging()
    
    docs_dir = find_docs_directory()
    
    handler = partial(SimpleHTTPRequestHandler, directory=docs_dir)
    
    try:
        httpd = HTTPServer(("", port), handler)
        server_url = f"http://localhost:{port}"
        logger.info(f"🚀 Documentation server started at {server_url}")
        logger.info(f"📖 Serving files from: {docs_dir}")
        logger.info("Press Ctrl+C to stop the server")
        
        def open_browser():
            webbrowser.open(server_url)
        
        threading.Thread(target=open_browser, daemon=True).start()
        
        httpd.serve_forever()
        
    except KeyboardInterrupt:
        logger.info("\n🛑 Documentation server stopped")
        httpd.shutdown()
        sys.exit(0)
    except Exception as e:
        logger.exception(f"Failed to start documentation server: {e}")
        sys.exit(1)

def main():
    logger = setup_logging()
    try:
        port = find_free_port()
        serve_documentation(port)
    except Exception as e:
        logger.exception(f"Failed to serve documentation: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()