import logging
import subprocess
import sys
from pathlib import Path
from typing import List

def setup_logging():
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S'
    )
    return logging.getLogger(__name__)

def find_proto_files(proto_root: str = "schema/proto/dart") -> List[str]:
    logger = setup_logging()
    
    cwd = Path.cwd().resolve()
    proto_root_path = (cwd / proto_root).resolve()
    
    if not proto_root_path.exists():
        logger.error(f"Directory not found: {proto_root_path}")
        sys.exit(1)
    
    files = list(proto_root_path.rglob("*.proto"))
    if not files:
        logger.warning(f"No .proto files found in {proto_root_path}")
        return []
    
    logger.info(f"Found {len(files)} proto files:")
    proto_files = []
    
    for f in files:
        try:
            proto_base = cwd / "schema/proto"
            relative_path = f.relative_to(proto_base)
            proto_files.append(str(relative_path))
            logger.info(f"  - {relative_path}")
        except ValueError as e:
            logger.error(f"Path calculation error for {f}: {e}")
            sys.exit(1)
    
    return sorted(proto_files, key=lambda x: (
        "service.proto" in x,
        x.count("/"),
        x
    ))

def run_command(command: List[str], logger: logging.Logger, capture_output: bool = False) -> subprocess.CompletedProcess:
    logger.info(f"Running command: {' '.join(command)}")
    
    try:
        result = subprocess.run(
            command,
            capture_output=capture_output,
            text=True,
            check=True
        )
        if capture_output:
            logger.debug(f"Command output: {result.stdout}")
        return result
    except subprocess.CalledProcessError as e:
        logger.error(f"Command failed with exit code {e.returncode}")
        if e.stdout:
            logger.error(f"STDOUT: {e.stdout}")
        if e.stderr:
            logger.error(f"STDERR: {e.stderr}")
        sys.exit(e.returncode)
    except Exception as e:
        logger.exception(f"Unexpected error executing command: {e}")
        sys.exit(1)

def ensure_directory_exists(path: str):
    Path(path).mkdir(parents=True, exist_ok=True)