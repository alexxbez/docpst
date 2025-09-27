#!/bin/python

import sys
import shutil
from pathlib import Path
import subprocess

def main():
    if len(sys.argv) == 3:
        if sys.argv[1] == 'new':
            create_directory(sys.argv[2])
        else:
            print("Usage: docs <command> <arg>")
            exit(1)
    elif len(sys.argv) == 2:
        if sys.argv[1] == 'watch':
            work_on_project()
        else:
            print("Usage: docs <command> <arg>")
            exit(1)
    else:
        print("Usage: docs <command> <arg>")
        exit(1)

def create_directory(name: str):
    directory = Path(name)
    try:
        directory.mkdir(parents=True, exist_ok=False) 
    except FileExistsError:
        print(f"Directory {name} already exists.")
        exit(1)

    template_dir = Path("/home/alexx/bin/docs/")

    try:
        shutil.copy(template_dir / "template.typ", directory / "main.typ")
        shutil.copytree(template_dir / "img", directory / "img")

    except FileNotFoundError as e:
        print(f"Template file not found: {e}")
        shutil.rmtree(directory)
        exit(1)

    except Exception as e:
            print(f"Error copying files: {e}")
            shutil.rmtree(directory)
            exit(1)

    print(f"Document {name} created.")

def work_on_project():
    if not Path("main.typ").exists():
        print("Not in a project directory. Run from project root.")
        return

    if not Path("main.pdf").exists():
        compile_process = subprocess.run(["typst", "compile", "main.typ"])

    zellij_open_process = subprocess.run(["zellij", "action", "new-tab", "-n", "main"])
    zellij_move_process = subprocess.run(["zellij", "action", "move-tab", "left"])
    
    typst_process = subprocess.Popen(["typst", "watch", "main.typ"])
    zathura_process = subprocess.Popen(["zathura", "main.pdf"])
    
    print("Watching for changes... Press Ctrl+C to stop")
    try:
        typst_process.wait()
    except KeyboardInterrupt:
        typst_process.terminate()
        zathura_process.terminate()
        zellij_close_process = subprocess.run(["zellij", "action", "close-tab"])
        print("Stopped watching")

if __name__ == "__main__":
    main()
