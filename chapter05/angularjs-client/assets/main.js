var myapp = angular.module('myapp', []);
myapp.controller('employeeController', function ($scope, $http) {
 $scope.addEmployee = function (id, firstName, lastName) {
     var data = {
       id: id,
       firstName: firstName,
       lastName: lastName
     };
   $http.post('/employee/add', data).then(function (response) {
   if (response.data)
       $scope.message = "Post Data Submitted Successfully!";
   });
 };
});
