'use strict';
const React = require('react');
var axios = require('axios');

import EmployeeList from './employee-list.jsx'
import AddEmployee from './add-employee.jsx'

export default class EmployeeApp extends React.Component {

	constructor(props) {
		super(props);
		this.state = {employees: []};
		this.addEmployee = this.addEmployee.bind(this);
		this.Axios = axios.create({
		    headers: {'content-type': 'application/json'}
		});
	}

	componentDidMount() {
		let _this = this;
		this.Axios.get('/employees')
		  .then(function (response) {
		    _this.setState({employees: response.data});
		  })
		  .catch(function (error) { });
	}

	addEmployee(employeeName){
		let _this = this;
		this.Axios.post('/employee/add', {
        	firstName: employeeName
         })
		  .then(function (response) {
		    _this.setState({employees: response.data});
		  })
		  .catch(function (error) { });
    }
	render() {
		return (
				<div>
				  <AddEmployee addEmployee={this.addEmployee}/>
				  <EmployeeList employees={this.state.employees}/>
		    </div>
			)
	}
}
