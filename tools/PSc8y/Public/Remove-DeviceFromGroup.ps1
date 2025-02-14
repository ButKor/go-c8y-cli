﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Remove-DeviceFromGroup {
<#
.SYNOPSIS
Unassign device from group

.DESCRIPTION
Unassign/delete a device from a group

.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/devicegroups_unassignDevice

.EXAMPLE
PS> Remove-DeviceFromGroup -Group $Group.id -ChildDevice $ChildDevice.id

Unassign a child device from its parent asset


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # Asset id (required)
        [Parameter(Mandatory = $true)]
        [object[]]
        $Group,

        # Child device (required)
        [Parameter(Mandatory = $true,
                   ValueFromPipeline=$true,
                   ValueFromPipelineByPropertyName=$true)]
        [object[]]
        $ChildDevice
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Delete"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "devicegroups unassignDevice"
        $ClientOptions = Get-ClientOutputOption $PSBoundParameters
        $TypeOptions = @{
            Type = ""
            ItemType = ""
            BoundParameters = $PSBoundParameters
        }
    }

    Process {

        if ($ClientOptions.ConvertToPS) {
            $ChildDevice `
            | Group-ClientRequests `
            | c8y devicegroups unassignDevice $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            $ChildDevice `
            | Group-ClientRequests `
            | c8y devicegroups unassignDevice $c8yargs
        }
        
    }

    End {}
}
