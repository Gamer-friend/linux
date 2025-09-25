# 🚀 Linux Docker Container Starter
This repository provides a ready-to-run Linux-based Docker container environment, optimized for quick setup in GitHub Codespaces. It exposes port 30099 for public access and includes a simple startup script and Docker Compose configuration.

# 🧰 Prerequisites
1. A GitHub account
2. Access to GitHub Codespaces
3. Basic familiarity with Docker and terminal commands

# ⚙️ Setup Instructions
1. Fork this repository to your GitHub account.
2. Open a Codespace from your forked repo.
3. Once the environment loads, open the terminal.
4. Make the startup script executable: (run these commands)

# Step 1
`chmod +x ./start.sh`
Make the Docker Compose file executable (optional but included for consistency):

# Step 2
`chmod +x ./docker-compose.yml`
Launch the container in detached mode:

# Step 3
`docker compose up -d`

# 🌐 Accessing the App
- The container exposes port 30099 publicly.
- You can access the running service via the Codespaces forwarded port interface.

📝 Notes
If you happen to have permission issues, please make sure Docker is running and your Codespace has enough privileges.
The start.sh script can be customized to include additional setup steps or service initialization.
