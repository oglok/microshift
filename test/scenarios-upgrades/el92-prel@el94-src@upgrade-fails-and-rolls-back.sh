#!/bin/bash

# Sourced from scenario.sh and uses functions defined there.

scenario_create_vms() {
    prepare_kickstart host1 kickstart.ks.template rhel-9.2-microshift-4."$(previous_minor_version)"
    launch_vm host1 rhel-9.2
}

scenario_remove_vms() {
    remove_vm host1
}

# CI will execute these tests on a set release branch, so source version will always be the latest code for that branch.
# I.e. CI jobs testing release-4.14 will always compile the latest 4.14 code as the source_version.
scenario_run_tests() {
    run_tests host1 \
        --variable "FAILING_REF:rhel-9.4-microshift-source" \
        --variable "REASON:fail_greenboot" \
        suites/upgrade/upgrade-fails-and-rolls-back.robot
}
