aws-stack-controller
====================

Application for starting up and shutting down a stack based on a tagging strategy

== How to build ==
Install go.
 go build
Test:
 % ./aws-stack-controller -h
 Usage of ./aws-stack-controller:
   -action="": 'startup' OR 'shutdown' without this, this does nothing
   -environment="": Name of environment to shutdown or startup
   -force=false: Force production actions on production tagged servers
   -publicKey="": aws iam publicKey
   -secretKey="": aws iam secretKey
   -stack="": Name of stack to shutdown or startup

== How to use ==
Ensure that all your instances are tagged with the following tags:
 Environment
 Stack
 StartOrder
 StopOrder
 Name

* StartOrder and StopOrder are numbers between 1 and N. (Doesn't N doesn't matter as long as it's above 1.)
* Environment should be one of: Test, UAT, Production, Dev (But not restricted to.)
* Stack is any name relevant to the stack.

When you call a startup or a shutdown, you need to specify the environment, and the stack name, it will group up the
instances based on these two values. Then it will act on them in the correct order, it will always start at 1 and
increment, regardless of operation.

When using the app, I recommend to create an IAM that can list, start and stop instances only, and use that with this
application.

So an example use:
 % ./aws-stack-controller -action="shutdown" -environment="UAT" -publicKey="" -privateKey="" -stack="CMS Stack"
Would shutdown the UAT CMS Stack.