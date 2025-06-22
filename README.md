# AI agent

This is a proof of concept for a working AI agent using the anthropic Go SDK in the terminal. It is using
Claude 3.5 Haiku. This model is hardcoded in the file `client/client.go` so it can be edited there to use another
model.

You need an anthropic api key to make this work locally. Set it up in your env like this:
```bash
export ANTHROPIC_API_KEY="<anthropic api key>"
```

## Example
![Example](https://github.com/DanielOchoa/ai-agent/blob/7a47ba27f571c82dadbe8da1a85c4f7b3514bb85/assets/example.png)

## Tools

It includes 3 tools so the AI can work with your file system:

- Read files
- List files
- Edit files

With these three included tools you can ask the LLM Model to check on files, explain to you contents of files,
edit or create files - adjust scripts for example, and more.

## Credits

Created by Daniel Ochoa based on the example by Thorsten Ball from  April 15, 2025 in the ampcode blog.

