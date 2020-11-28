# 1. Names should be Intention-Revealing!
Names should be increase intention in the code.. beacuse variables doesn't just hold a value, it should show it's value in it's name. 
It should answer questions such as What it it? Why is it here? or What does it do? how it's used and etc...
<br />
Look at the code blow and try to inderstand what is it..
```
```

**Variable should reveal all information about them self in their name. They do not created for nothing just hold a simple value**

for example do not refer to a group of acounts as _accountlist_ unless it's acctualy a _list_ ! you can refer it a _acounts_ or _account\_group_

# 2. Make meaningful distinctions

### 1.avoid number series naming
you think you have two variabl3  hold somehow the same thing then it's not proper name number after them( like a1, a2)
for example:


```
def copy_chars(a1, a2):
    for i in range(len(a2)):
        a1[i] = a2[i]
```

you can see both a1 and a2 hold char but it this code redeable? can you understand easily? you can name a1 as _destination_ and a2 as _source_


### 2. avoid noise words

*	imaging have a _company_ variable , you think you have another variables _productinfo_ and_productdata_, both info and data represent the same thing

*	you can not put a table variable for a table and variable for a variable

*	you can not have _nameString_ beacuse names are string can they be float? or UserObject. users are object them self

*	Imagine finding one class named Customer and another named CustomerObject . What should you understand as the distinction? Which one will represent the best path to a customer’s payment history?


**The programmer should easily distinguish tow vaiable, for example in the absents of conventions what is the diffrent between _moneyAmount_ and money, or _customer_ and _customerInfo_?**


### 3. Use Pronounceable Names

beacuse programming is a social activity and humans can remmember pronounceable names easily it's better to put your variable name if a way it can be pronounce easy for example:

`genymdhms` what does it means? it's a generation date: year, month, day, hours, minute, second) how can you pronounce it? you change this variable name to `generationTimestamp`


### 4.Use Searchable Names
It's improtant to use searchable names. like you have 7 weeks of a day, it's better to put it in a variable like DAYS_OF_WEEK and then use it all over the code


## CLASS NAMES

class names should be **NOUN** like _Customer_ or _WikiPage_ avoid works like _Manager_ or _Processor_ beacuse class names should be noun

## Method names
Method names should be **VERB** or **VERB PHRASE NAMES** like _postPayment_ or _deletePage_


# 3- Don't be cute!

