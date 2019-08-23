## Clean code series

* On encapsulation: for business objects, we want to encapsulate the fields while we can make the fields of data 
objects. See `User`, `Codecast` and `License` vs `PresentableCodecast`
* We don't want the Presenters to know about the business objects. That's why we don't pass the `Codecast` object as an 
argument to `PresentableCodecasts`.
* The question when deciding whether to make create a subclass or go with type flags is to ask if we are dealing with 
data or behavior. If it is data we want type codes; if it is behavior we want derivatives