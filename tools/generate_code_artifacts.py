#!/usr/bin/env python3
import logging
import sys
from pathlib import Path

from tools.common import setup_logging, find_proto_files, run_command, ensure_directory_exists

def generate_protos(proto_files: list):
    logger = setup_logging()
    
    if not proto_files:
        logger.info("No proto files to process, skipping generation")
        return
    
    ensure_directory_exists("pkg")
    ensure_directory_exists("docs")
    
    cmd = [
        "protoc",
        "--proto_path=schema/proto",
        "--proto_path=schema/proto/third_party",
        "--go_out=paths=source_relative:pkg",
        "--go-grpc_out=paths=source_relative:pkg",
        "--grpc-gateway_out=paths=source_relative:pkg",
        "--grpc-gateway_opt=logtostderr=true",
        "--grpc-gateway_opt=generate_unbound_methods=true",
        "--validate_out=lang=go,paths=source_relative:pkg",
        "--openapiv2_out=docs",
        "--openapiv2_opt=logtostderr=true",
        "--openapiv2_opt=allow_merge=true",
        "--openapiv2_opt=merge_file_name=api"
    ] + proto_files
    
    run_command(cmd, logger)
    logger.info("✅ Code artifacts generated successfully")

def main():
    logger = setup_logging()
    try:
        proto_files = find_proto_files()
        generate_protos(proto_files)
    except Exception as e:
        logger.exception(f"Failed to generate code artifacts: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()