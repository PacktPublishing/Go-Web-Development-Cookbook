var vue_det = new Vue({
   el: '#form',
   data: {
      message: 'Employee Dashboard',
      id: '',
      firstName:'',
      lastName:''
   },
   methods: {
     addEmployee: function() {
       this.$http.post('/employee/add', {
             id: this.id,
             firstName:this.firstName,
             lastName:this.lastName
           }).then(response => {
              console.log(response);
            }, error => {
              console.error(error);
            });
     }}
});
