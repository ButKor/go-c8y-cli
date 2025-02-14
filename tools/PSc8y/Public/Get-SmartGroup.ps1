﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Get-SmartGroup {
<#
.SYNOPSIS
Get smart group

.DESCRIPTION
Get an smart group

.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/smartgroups_get

.EXAMPLE
PS> Get-SmartGroup -Id $smartgroup.id

Get smart group by id

.EXAMPLE
PS> Get-SmartGroup -Id $smartgroup.name

Get smart group by name


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # Smart group ID (required)
        [Parameter(Mandatory = $true,
                   ValueFromPipeline=$true,
                   ValueFromPipelineByPropertyName=$true)]
        [object[]]
        $Id
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Get"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "smartgroups get"
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
            | c8y smartgroups get $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            $Id `
            | Group-ClientRequests `
            | c8y smartgroups get $c8yargs
        }
        
    }

    End {}
}
