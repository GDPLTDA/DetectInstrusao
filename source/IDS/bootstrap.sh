#update inicial
sudo apt-get update

#basico

sudo apt-get install -y vim 2> /dev/null

sudo apt-get install -y apache2 2> /dev/null

# golang
sudo apt-get install -y python-software-properties 2> /dev/null
sudo add-apt-repository ppa:duh/golang
sudo apt-get update
sudo apt-get install -y golang 2> /dev/null