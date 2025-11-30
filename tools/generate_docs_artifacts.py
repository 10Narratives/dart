import logging
import sys
from pathlib import Path
import json

from tools.common import setup_logging, run_command, ensure_directory_exists

def find_swagger_files(docs_root: str = "docs") -> list:
    logger = setup_logging()
    
    cwd = Path.cwd().resolve()
    docs_root_path = (cwd / docs_root).resolve()
    
    if not docs_root_path.exists():
        logger.error(f"Documentation root not found: {docs_root_path}")
        sys.exit(1)
    
    swagger_files = list(docs_root_path.rglob("*.swagger.json"))
    if not swagger_files:
        logger.warning("No swagger files found. Generating code artifacts first...")
        from tools.generate_code_artifacts import main as generate_code
        generate_code()
        swagger_files = list(docs_root_path.rglob("*.swagger.json"))
    
    if not swagger_files:
        logger.error("No swagger files found after code generation")
        sys.exit(1)
    
    logger.info(f"Found {len(swagger_files)} swagger files:")
    relative_files = []
    
    for f in swagger_files:
        try:
            relative_path = f.relative_to(cwd)
            relative_files.append(str(relative_path))
            logger.info(f"  - {relative_path}")
        except ValueError as e:
            logger.error(f"Path calculation error for {f}: {e}")
            sys.exit(1)
    
    return relative_files

def merge_swagger_files(swagger_files: list, output_file: Path):
    logger = setup_logging()
    
    if not swagger_files:
        logger.warning("No files to merge")
        return
    
    first_file = Path(swagger_files[0])
    with open(first_file, 'r') as f:
        merged = json.load(f)
    
    for file_path in swagger_files[1:]:
        file_path = Path(file_path)
        with open(file_path, 'r') as f:
            data = json.load(f)
        
        if 'paths' in data:
            if 'paths' not in merged:
                merged['paths'] = {}
            for path, methods in data['paths'].items():
                if path not in merged['paths']:
                    merged['paths'][path] = methods
                else:
                    merged['paths'][path].update(methods)
        
        if 'definitions' in data:
            if 'definitions' not in merged:
                merged['definitions'] = {}
            for definition, schema in data['definitions'].items():
                if definition not in merged['definitions']:
                    merged['definitions'][definition] = schema
    
    output_file.parent.mkdir(parents=True, exist_ok=True)
    
    with open(output_file, 'w') as f:
        json.dump(merged, f, indent=2)
    
    logger.info(f"✅ Merged {len(swagger_files)} swagger files into {output_file}")

def generate_html_docs():
    logger = setup_logging()
    
    swagger_files = find_swagger_files()
    
    html_dir = Path("docs/html")
    ensure_directory_exists(str(html_dir))
    
    merged_swagger = html_dir / "api.swagger.json"
    merge_swagger_files(swagger_files, merged_swagger)
    
    output_file = html_dir / "index.html"
    
    try:
        run_command([
            "npx", "redoc-cli", "bundle",
            str(merged_swagger),
            "--output", str(output_file),
            "--title", "Dart API Documentation"
        ], logger)
        logger.info(f"✅ HTML documentation generated at {output_file}")
        
        with open(html_dir / "version.json", 'w') as f:
            import datetime
            json.dump({
                "generated_at": datetime.datetime.now().isoformat(),
                "swagger_files": [str(f) for f in swagger_files],
                "version": "1.0.0"
            }, f, indent=2)
        
    except Exception as e:
        logger.warning(f"Failed to generate HTML docs with redoc-cli: {e}")
        logger.warning("Try installing redoc-cli: npm install -g redoc-cli")
        with open(output_file, 'w') as f:
            f.write(f"""<!DOCTYPE html>
<html>
<head>
    <title>API Documentation</title>
    <style>body {{ font-family: Arial, sans-serif; padding: 20px; }}</style>
</head>
<body>
    <h1>API Documentation</h1>
    <p>Failed to generate documentation: {str(e)}</p>
    <p>Install redoc-cli: <code>npm install -g redoc-cli</code></p>
</body>
</html>""")
        logger.info(f"Created placeholder documentation at {output_file}")

def main():
    logger = setup_logging()
    try:
        generate_html_docs()
    except Exception as e:
        logger.exception(f"Failed to generate documentation artifacts: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()