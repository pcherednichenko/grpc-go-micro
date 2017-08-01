this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'hello_services_pb'

def main
  stub = Go::Micro::Srv::Greeter::Say::Stub.new('localhost:58941', :this_channel_is_insecure)
  user = ARGV.size > 0 ?  ARGV[0] : 'Cryptopay Ruby'
  message = stub.hello(Go::Micro::Srv::Greeter::Request.new(name: user)).msg
  p "Greeting: #{message}"
end

main