﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Get-FirmwareVersion {
<#
.SYNOPSIS
Get firmware package version

.DESCRIPTION
Get an existing firmware package version

.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/firmware_versions_get

.EXAMPLE
PS> Get-FirmwareVersion -Id $mo.id

Get a firmware package

.EXAMPLE
PS> Get-ManagedObject -Id $mo.id | Get-FirmwareVersion

Get a firmware package (using pipeline)


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # Firmware Package version id or name (required)
        [Parameter(Mandatory = $true,
                   ValueFromPipeline=$true,
                   ValueFromPipelineByPropertyName=$true)]
        [object[]]
        $Id,

        # Firmware package id or name (used to help completion be more accurate)
        [Parameter()]
        [object[]]
        $Firmware,

        # Include parent references
        [Parameter()]
        [switch]
        $WithParents
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Get"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "firmware versions get"
        $ClientOptions = Get-ClientOutputOption $PSBoundParameters
        $TypeOptions = @{
            Type = "application/vnd.com.nsn.cumulocity.managedObject+json"
            ItemType = ""
            BoundParameters = $PSBoundParameters
        }
    }

    Process {

        if ($ClientOptions.ConvertToPS) {
            $Id `
            | Group-ClientRequests `
            | c8y firmware versions get $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            $Id `
            | Group-ClientRequests `
            | c8y firmware versions get $c8yargs
        }
        
    }

    End {}
}
