ssh_access
=========

Role setups a host to have access under sys_ci user

Requirements
------------

See requirements.txt

Role Variables
--------------

Low level variables:
sshd_name - Required. Name of sshd server, default: 'ssh'
sys_ci_user - Required. Name of system user to create, default: 'sys_ci'

High level variables:
id_rsa - public ssh-key to copy to {{ sys_ci_user }} home directory


Example Playbook
----------------

---
- name: Setup sys_ci ssh-access
  hosts: all
  roles:
    - role: ssh_access
      tags: ssh
  become: true

License
-------

Apache 2.0

Author Information
------------------

Anton Shadrin aashadrin@edu.hse.ru