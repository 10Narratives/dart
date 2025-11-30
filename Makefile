generate:
	PYTHONPATH=. python3 tools/generate_code_artifacts.py
	PYTHONPATH=. python3 tools/generate_docs_artifacts.py

docs: generate
	PYTHONPATH=. python3 tools/serve_docs.py

clean:
	PYTHONPATH=. python3 tools/clean_generated_artifacts.py

.PHONY: generate docs clean