local box = require("box")

local function bootstrap()
	box.schema.user.grant('guest', 'read,write,execute', 'universe')
	
	fspace = box.schema.space.create('fingerprintSpace')
	fspace:format(
		{{name='id', type='unsigned'},
		{name='fingerprint', type='string'}}
	)
	fspace:create_index('primary', {type = 'tree', parts = {'id'}})
	fspace:create_index('fingerprint', {unique = true, type = 'tree', parts = {'fingerprint'}})
	-- fspace:create_index('clientID', {unique = false, type = 'tree', parts = {'clientID'}})

	cnspace = box.schema.space.create('chatroomNames')
	cnspace:format(
		{
			{name='chatroomId', type='unsigned'},
			{name='chatroomName', type='string'}
		})

	cnspace:create_index('general', {type = 'hash', parts = {'chatroomId'}})

end

box.once('onstartup', bootstrap)
--[[ return {
	bootstrap = bootstrap;
} ]]