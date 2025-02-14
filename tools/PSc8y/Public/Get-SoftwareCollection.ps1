﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Get-SoftwareCollection {
<#
.SYNOPSIS
Get software collection

.DESCRIPTION
Get a collection of software packages (managedObjects) based on filter parameters

.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/software_list

.EXAMPLE
PS> Get-SoftwareCollection

Get a list of software packages


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # Software name filter
        [Parameter()]
        [string]
        $Name,

        # Software description filter
        [Parameter()]
        [string]
        $Description,

        # Software device type filter
        [Parameter()]
        [string]
        $DeviceType
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Get", "Collection"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "software list"
        $ClientOptions = Get-ClientOutputOption $PSBoundParameters
        $TypeOptions = @{
            Type = "application/vnd.com.nsn.cumulocity.managedObjectCollection+json"
            ItemType = "application/vnd.com.nsn.cumulocity.managedObject+json"
            BoundParameters = $PSBoundParameters
        }
    }

    Process {

        if ($ClientOptions.ConvertToPS) {
            c8y software list $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            c8y software list $c8yargs
        }
    }

    End {}
}
