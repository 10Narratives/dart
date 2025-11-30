#!/usr/bin/env python3
import logging
import sys
from pathlib import Path
from typing import List

from tools.common import setup_logging

def clean_artifacts(patterns: List[str], logger: logging.Logger):
    deleted_count = 0
    
    for pattern in patterns:
        for file_path in Path(".").glob(pattern):
            if file_path.is_file():
                try:
                    file_path.unlink()
                    logger.info(f"Deleted: {file_path}")
                    deleted_count += 1
                except Exception as e:
                    logger.error(f"Failed to delete {file_path}: {e}")
    
    return deleted_count

def main():
    logger = setup_logging()
    
    try:
        patterns_to_clean = [
            "pkg/**/*.pb.go",
            "pkg/**/*.grpc.pb.go",
            "pkg/**/*.pb.gw.go",
            "pkg/**/*.pb.validate.go",
            "docs/**/*.swagger.json",
            "docs/**/*.swagger.yaml",
            "docs/**/index.html",
            "docs/**/api.html",
            "docs/**/version.json",
            "docs/**/redoc.standalone.js",
            "*.tmp",
            "*.log",
            "docs/dart/**/*.json",
            "docs/dart/**/*.yaml",
            "docs/html/**/*",
            "docs/api/**/*"
        ]
        
        logger.info("🧹 Cleaning generated artifacts...")
        
        existing_dirs = set()
        for pattern in patterns_to_clean:
            for file_path in Path(".").glob(pattern):
                if file_path.is_file():
                    existing_dirs.add(str(file_path.parent))
        
        deleted_count = clean_artifacts(patterns_to_clean, logger)
        
        for dir_path in sorted(existing_dirs, key=len, reverse=True):
            dir_obj = Path(dir_path)
            if dir_obj.exists() and not any(dir_obj.iterdir()):
                try:
                    dir_obj.rmdir()
                    logger.info(f"Removed empty directory: {dir_path}")
                except Exception as e:
                    logger.debug(f"Could not remove directory {dir_path}: {e}")
        
        logger.info(f"✅ Successfully cleaned {deleted_count} files")
        
    except Exception as e:
        logger.exception(f"Cleanup failed: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()