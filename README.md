# app build step
chmod u+x build.sh
./build.sh

# CentOS7
sudo systemctl stop firewalld
sudo systemctl mask firewalld

# url
http://192.168.80.10:8000/
