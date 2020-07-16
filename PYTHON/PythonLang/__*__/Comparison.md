# What this file is?

this file is brief defenition for following attribures

*	`__eq__`[==]
*	`__lt__`[<]
*	`__le__`[<=]
*	`__gt__`[>]
*	`__ge__`[>=]


this are used when we want to make compersion

like class base compersion

eg
```
class car:
	def __init__(self, name, mile):
		self.name = name
		self.mile = mile
	def __eq__(self, other):
		return self.mile == other.mile
	def __gt__(self, other):
		return self.mile > other.mile


if __name__ == '__main__':
	car1 = car('pride', 3444)
	car2 = car('persia', 1233)

	car1 == car2


	# We will face error becuse we do not specified __lt__
	car1 < car2
```

