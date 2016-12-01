#!/bin/bash

(
    $starting_step "Create the public key and setup ssh config"
    [[ -f ${NODE_DIR}/root@${vm_name} ]]
    $skip_step_if_already_done; set -ex
    ssh-keygen -t rsa -N "" -f ${NODE_DIR}/root@${vm_name}
    chmod 600 ${NODE_DIR}/root@${vm_name}

) ; prev_cmd_failed

(
    $starting_step "Install authorized ssh key for ${vm_name}"
    sudo chroot ${TMP_ROOT} /bin/bash -c \
         "[[ -f /root/.ssh/authorized_keys ]]"
    $skip_step_if_already_done; set -xe

sudo mkdir -p -m 700 ${TMP_ROOT}/root/.ssh
sudo chmod 700 ${TMP_ROOT}/root/.ssh
sudo mv ${NODE_DIR}/root@${vm_name}.pub ${TMP_ROOT}/root/.ssh/authorized_keys
sudo chmod 600 ${TMP_ROOT}/root/.ssh/authorized_keys

sudo chroot ${TMP_ROOT} /bin/bash -ex <<EOS
sed -i \
-e 's,^PermitRootLogin .*,PermitRootLogin yes,' \
-e 's,^PasswordAuthentication .*,PasswordAuthentication yes,' \
-e 's,^DenyUsers.root,#DenyUsers root,' \
\
/etc/ssh/sshd_config
EOS
  
) ; prev_cmd_failed

if [[ ${vna} == true ]]; then
    (
        $starting_step "Install autorized key for jenkins"
        [[ $(sudo wc -l ${TMP_ROOT} /root/.ssh/authorized_keys | awk '{ print $1 }') -eq 2 ]]
        $skip_step_if_already_done; set -xe
        sudo chroot ${TMP_ROOT} /bin/bash -ex <<EOS
cat <<CFG >> /root/.ssh/authorized_keys
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAtN9MMMPtRMY9wGRXvivF5mkSfY+/ZN7wfLPu1JMtCgxN9TgEmc9ag95e+EcWke4e6qiYFFfarU7OtMSIu4qKjXnolbcjzbm+dHBrQKsK4rZEggOtORANcsZRWGYhE78fvGzR4Bpfs3Gy0ko1BLEOATBPTreai2T5168vKrJG0tdA/77pvAwi41kppqQXlULuU3R20ynWZRsrX8JPb9BJIz/jusVskDX5U3Waw3jUi2qSfi+gBaVKb3uFk/ctUXec/etRWut1oiodn/Gd8WvrlWFuG/0Ob0aqg8bO51Yl657Kf3llFNvKYUPOfk0g8VvwHGWtEDjQ28/6RteQRiPMEw== root@kemumaki
CFG
EOS
    ) ; prev_cmd_failed
fi
