#!/usr/bin/env tarantool

local box = require("box")
box.cfg{
  listen = 3301
}	

local boot = require("onstartup")