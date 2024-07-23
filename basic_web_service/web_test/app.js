//const baseURl = 'https://api.sampleapis.com/coffee/hot';
/*
const baseURl = 'http://localhost:5000/course';
fetch(baseURl)
    .then(resp => resp.json())
    .then(data => display(data));

function display(data) {
    document.querySelector("pre").innerHTML = JSON.stringify(data, null, 2);
}
*/
fetch('http://localhost:5000/course')
.then(function(response) {
    return response.json();
})
.then(function(data) {
    appendData(data);
})
.catch(function(err) {
    console.log('error: ' + err);
});
function appendData(data) {
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++) {
        var div = document.createElement("div");
        div.innerHTML = 'CourseID: ' + data[i].Id + ' ' + data[i].Name + ' ' + data[i].Price + ' ' + data[i].Instructor
        mainContainer.appendChild(div);
    }
}