const React = require('react');

export default class Employee extends React.Component{
	render() {
		return (
			<tr>
				<td>{this.props.employee.firstName}</td>
			</tr>
		)
	}
}
