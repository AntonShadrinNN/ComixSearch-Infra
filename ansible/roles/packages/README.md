pacakges
=========

Role installs packages on host under Debian-based distributions

Requirements
------------

See requirements.txt

Role Variables
--------------

Low level variables:
main_packages - Required. List of packages to be installed on host, default: 'golang-1.21-go'
additional_packages - Additional packages, default: 'vim'

Example Playbook
----------------

---
- name: Setup sys_ci ssh-access
  hosts: all
  roles:
    - role: packages
      tags: packages
  become: true

License
-------

Apache 2.0

Author Information
------------------

Anton Shadrin aashadrin@edu.hse.ru