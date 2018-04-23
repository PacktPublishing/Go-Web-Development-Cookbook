<!DOCTYPE html>
<html>
   <body>
      <table border= "1" style="width:100%;">
         {{range .employees}}
         <tr>
            <td>{{.Id}}</td>
            <td>{{.FirstName}}</td>
            <td>{{.LastName}}</td>
         </tr>
         {{end}}
      </table>
   </body>
</html>
