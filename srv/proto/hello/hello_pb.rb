# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "go.micro.srv.greeter.Request" do
    optional :name, :string, 1
  end
  add_message "go.micro.srv.greeter.Response" do
    optional :msg, :string, 1
  end
end

module Go
  module Micro
    module Srv
      module Greeter
        Request = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.srv.greeter.Request").msgclass
        Response = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.srv.greeter.Response").msgclass
      end
    end
  end
end
