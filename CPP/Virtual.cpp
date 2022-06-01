#include<iostream>
#include<typeinfo>
using namespace std;


class A{
	public:
		void show(){
			cout << typeid(this).name();
		}
};

class B : public virtual A {
};
  
class C : public virtual  A {
};
  
class D : public B, public C {
};

int main(){
   D tester;
   tester.show();
}
