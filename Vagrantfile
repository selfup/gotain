Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.network "forwarded_port", guest: 80, host: 8080

  config.vm.provider 'virtualbox' do |vb|
    vb.name = 'gotain'
    vb.cpus = 1
    vb.memory = 1024
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]
  end
end
