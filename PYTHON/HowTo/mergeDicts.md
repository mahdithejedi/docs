# How to merge two python dictionries?

You think you have three dictionries:
<br />
```
d1={1:2,3:4}
d2={5:6,7:9}
d3={10:8,13:22}
```
<br />
and you want to have something like this:
<br />
```
d4={1:2,3:4,5:6,7:9,10:8,13:22}
```

## First method(Fastest) : exploit the dict constructor to the hilt, then one update:
```
d1={1:2,3:4}
d2={5:6,7:9}
d3={10:8,13:22}'
d4 = dict(d1, **d2)
d4.update(d3)
```

## Second method: a loop of update calls on an initially-empty dict:
```
d1={1:2,3:4}
d2={5:6,7:9}
d3={10:8,13:22}
d4 = {}
for d in (d1, d2, d3): 
	d4.update(d)
```

## Third method: one copy-ctor and two updates

```
d1={1:2,3:4}
d2={5:6,7:9}
d3={10:8,13:22}
d4 = dict(d1)
for d in (d2, d3): 
	d4.update(d)
```


## 4th method (slowest): oncatenate the items and call dict on the resulting list
```
d1={1:2,3:4}
d2={5:6,7:9}
d3={10:8,13:22}
d4 = dict(d1.items() + d2.items() + d3.items())
```
<br />
***IMPORTANT POINT: this does not work in python 3***


## 6TH & NEW METHOD: introduce in python 3.9.0a4 
```
d4 = (d1 | d2) | d3
```
<br />
***IMPORTANT NOTE: this work just in python 3.9.0a4 and later***


## 7TH method: Ugly boy!
```
d4 = {**d1, **d2, **d3}
```

### Resources
[pep-584](https://www.python.org/dev/peps/pep-0584/#d1-d2)
<br />
[pep-448](https://www.python.org/dev/peps/pep-0448/)
<br />
[stackoverflow](https://stackoverflow.com/a/1784128/9651641)
<br />
