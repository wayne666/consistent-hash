# consistent hash

-----------------

## Description

	Simple ConsistentHash library

## Installation

	$ go get github.com/WayneZhouChina/consistent-hash

## Usage
	import "github.com/WayneZhouChina/consistent-hash"

	c := ConsistentHash.New()
	c.Add("stationA")
	c.Add("stationB")
	c.Add("stationC")
	c.Add("stationD")

	station, error := c.Get("Wayne")

## AUTHOR

	Written by WayneZhou, cumtxhzyy[at]gmail.com

## COPYRIGHT

	Copyright (c) 2015 WayneZhou. This library is free software; you can redistribute it and/or modify it.
