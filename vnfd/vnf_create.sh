#!bin/bash

#openstack vnf create --vnfd-name mongo mongo --vim-name chin-vim --description "mongodb"
openstack vnf create --vnfd-name upf1 upf1 --vim-name chin-vim --description "upf1"
openstack vnf create --vnfd-name upf2 upf2 --vim-name chin-vim --description "upf2"
openstack vnf create --vnfd-name upf3 upf3 --vim-name chin-vim --description "upf3"

sleep 10

openstack vnf create --vnfd-name nrf nrf --vim-name chin-vim --description "nrf"
sleep 10
openstack vnf create --vnfd-name amf amf --vim-name chin-vim --description "amf"
sleep 10
openstack vnf create --vnfd-name smf smf --vim-name chin-vim --description "smf"
sleep 10
openstack vnf create --vnfd-name udr udr --vim-name chin-vim --description "udr"
sleep 10
openstack vnf create --vnfd-name pcf pcf --vim-name chin-vim --description "pcf"
sleep 10
openstack vnf create --vnfd-name udm udm --vim-name chin-vim --description "udm"
sleep 10
openstack vnf create --vnfd-name nssf nssf --vim-name chin-vim --description "nssf"
sleep 10
openstack vnf create --vnfd-name ausf ausf --vim-name chin-vim --description "ausf"
#openstack vnf create --vnfd-name basic_setup basic_setup --vim-name chin-vim --description "basic_setup"
