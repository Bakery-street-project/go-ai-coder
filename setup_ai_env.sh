#!/usr/bin/env bash

# ==============================================================================
# Comprehensive Ubuntu AI Development Environment Setup Script
#
# This script installs NVIDIA drivers, Docker, the NVIDIA Container Toolkit,
# and Zsh/Oh-My-Zsh for a complete GPU-accelerated development setup.
# ==============================================================================

# Exit immediately if a command exits with a non-zero status.
set -e

# --- Color Definitions for Output ---
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting the Ultimate Ubuntu AI Development Environment Setup...${NC}"

# --- STEP 1: System Update and Essential Tools ---
echo -e "\n${YELLOW}---> Step 1: Updating system packages and installing essentials...${NC}"
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install -y curl wget git zsh ca-certificates gnupg 

# --- STEP 2: Install Oh My Zsh for an Enhanced Terminal ---
echo -e "\n${YELLOW}---> Step 2: Installing Oh My Zsh...${NC}"
if [ -d "$HOME/.oh-my-zsh" ]; then
  echo "Oh My Zsh is already installed. Skipping."
else
  # The --unattended flag prevents prompts and shell changes .
  sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
  # Change the default shell to Zsh non-interactively .
  sudo chsh -s "$(which zsh)" "${USER}"
  echo -e "${GREEN}Oh My Zsh installed and set as default shell.${NC}"
  echo -e "${YELLOW}You must log out and log back in for the shell change to take effect.${NC}"
fi

# --- STEP 3: Install NVIDIA GPU Drivers ---
echo -e "\n${YELLOW}---> Step 3: Installing recommended NVIDIA drivers...${NC}"
# This command detects the GPU and installs the best driver .
sudo ubuntu-drivers autoinstall
echo -e "${GREEN}NVIDIA driver installation command issued.${NC}"
echo -e "${YELLOW}A REBOOT IS REQUIRED after this script finishes to load the new drivers.${NC}"

# --- STEP 4: Install Docker Engine ---
echo -e "\n${YELLOW}---> Step 4: Installing Docker Engine...${NC}"
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin 
# Add your user to the docker group to run docker commands without sudo .
sudo usermod -aG docker ${USER}
echo -e "${GREEN}Docker has been installed.${NC}"
echo -e "${YELLOW}You need to log out and log back in for Docker group changes to take effect!${NC}"

# --- STEP 5: Install NVIDIA Container Toolkit ---
echo -e "\n${YELLOW}---> Step 5: Installing the NVIDIA Container Toolkit...${NC}"
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \
  && curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
    sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
    sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list

sudo apt-get update
sudo apt-get install -y nvidia-container-toolkit 
echo -e "${GREEN}NVIDIA Container Toolkit has been installed.${NC}"

# --- STEP 6: Configure Docker and Final Instructions ---
echo -e "\n${YELLOW}---> Step 6: Configuring Docker to use the NVIDIA runtime...${NC}"
# This command makes Docker aware of the NVIDIA runtime and restarts Docker .
sudo nvidia-ctk runtime configure --runtime=docker
sudo systemctl restart docker
echo -e "${GREEN}Docker runtime configured and restarted.${NC}"

echo -e "\n\n${GREEN}====================================================="
echo -e "            SETUP COMPLETE!"
echo -e "=====================================================${NC}"
echo -e "\n${YELLOW}CRITICAL NEXT STEPS:${NC}"
echo -e "1. You ${YELLOW}MUST REBOOT${NC} your computer now to load the NVIDIA drivers."
echo -e "   Run: ${GREEN}sudo reboot${NC}"
echo -e "2. After rebooting, open a new terminal to verify the setup."
echo -e "3. Verify the driver: run ${GREEN}nvidia-smi${NC}. It should show your GPU details ."
echo -e "4. Verify Docker's GPU access: run the following command:"
echo -e "   ${GREEN}docker run --rm --gpus all nvidia/cuda:12.4.1-base-ubuntu22.04 nvidia-smi${NC} "
echo -e "   The output should match the output from 'nvidia-smi'."
echo -e "\nHappy coding!"

