
#!/bin/bash
cd $(dirname $0)

# Get repopull binary
sudo mkdir /usr/local/bin
sudo cp ./repopull /usr/local/bin
sudo chmod +x /usr/local/bin/repopull

# Build systemd service
sudo touch /etc/systemd/system/repopull.service
echo "[Unit]" > /etc/systemd/system/repopull.service
echo "Description=repopull service" >> /etc/systemd/system/repopull.service
echo "After=network.target" >> /etc/systemd/system/repopull.service
echo "[Service]" >> /etc/systemd/system/repopull.service
echo "ExecStart=/usr/local/bin/repopull" >> /etc/systemd/system/repopull.service
echo "[Install]" >> /etc/systemd/system/repopull.service
echo "WantedBy=multi-user.target" >> /etc/systemd/system/repopull.service

# Enable the repopull service
sudo systemctl enable repopull.service && sudo systemctl start repopull.service
sudo systemctl status repopull.service
