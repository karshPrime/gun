# gun - Go rUN

**gun** is a CLI utility to unify build, run, debug, and test workflows across multiple programming
languages. It uses a simple config system to define how a project should compile and run, making it
easier to manage different types of codebases with one consistent command.

This means whether you’re working on a system C project, an embedded C project, a Python script, or
something else, you don’t need to remember language-specific build/run steps — just define it once
and run it the same way every time.

---

## Features

- Run, build, test, clean, and debug from a single command
- Local and global config system
- Templating support for project scaffolding
- Built-in license manager
- Git setup on project creation
- Simple and extendable `.toml` config format

---

⚠️ Work In Progress  ⚠️

### Current Progress
- [x] Set project structure
- [x] Global config parser
- [x] Local config parser
- [x] Fallback to global config when local is missing
- [x] `run`, `compile`, and `debug` commands implemented
- [x] Git repository initialisation
- [x] Help command with detailed usage and logging
- [x] Parse `init` config definitions
- [x] Create new project with structure and files
- [x] Copy license file into project
- [x] Copy template files into project
- [x] **`init` command: Completed**
- [x] View/manage saved licenses
- [x] Replace current license
- [x] **License system: Completed**
- [ ] View/manage saved templates
- [ ] Create new template files
- [ ] Add templates to projects
- [ ] **Template system: Completed**

