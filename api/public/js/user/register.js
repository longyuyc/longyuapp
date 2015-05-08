/**
 * Created by ThinkPad on 2015/5/7.
 */
function Register($scope,$http){
    $scope.comfRegister=function(){
     $http({
         method:'POST',
         url:'',
         data:'',
         headers:{'Content-Type':'application/x-www-form-urlencoded'}
     }).success(function(data){
         alert(data)
     }).error(function(error){
         alert(error)
     });
    };
};