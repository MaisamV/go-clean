Every change or addition in this project must follow the architecture guidelines and rules. If you don't remember the architecture guidelines and rules, You **MUST** read `docs/architecture.md` file.

To implement a new technology that needs to import new external library or framework in the project You **MUST** read `docs/tech_stack.md` to find the specified tech and follow it's guidelines and rules. If no tech was defined in the doc, ask the user with suggestions that you recommend and let the user choose and update the `docs/tech_stack.md`.

Whenever you need to implement a new feature or add a new module, You **MUST** be aware of current features and if you don't remember even a part of `docs/features.md` you **MUST** read it to have a good understanding of project features. After changing or implementing a feature you **MUST** update `docs/features.md` doc to reflect current features implementation.

After enough meaningful changes or development, create a commit with a one-line commit indicating to most important change/s, the branching, branch name and commit messages should follow gitflow rules.

to run the app you should use `docker-compose up --build -d` command.

If you want to build an output binary file from the project, you should create the output in ./bin folder.

if you want to use `git diff` command, you **MUST** use it with `--no-pager` option.