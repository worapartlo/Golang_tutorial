fetch('http://localhost:8000/api/cargo')
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        dropdownData(data);
    })
    .catch(function (err) {
        console.log('error: ' + err);
    });
/*
function appendData(data) {
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++) {
        var div = document.createElement("div");
        div.innerHTML = 'CargoID: ' + data[i].CargoID + ' ' + data[i].CargoName + ' ' + data[i].Amount + ' ' + data[i].CargoDetail
        mainContainer.appendChild(div);
    }
}*/

function dropdownData(data) {
    var dropdown = document.getElementById('dropdownlist');

    for (var i = 0; i < data.length; i++) {
        var option = document.createElement('option');
        option.value = data[i].CargoID;
        option.text = data[i].CargoName;
        dropdown.appendChild(option);
    }
    dropdown.addEventListener('change', function(){
        displayData(data, this.value);
    });
}

function displayData(data, selectedCargo){
    var cargoData = document.getElementById('cargoDetail');
    cargoData.innerHTML = '';
    var selectCargo = data.find(item => item.CargoID == selectedCargo);
    if (selectCargo) {
        var div = document.createElement("div");
        div.innerHTML = 'CargoID: ' + selectCargo.CargoID + '<br>' +
                        'CargoName: ' + selectCargo.CargoName + '<br>' +
                        'Amount: ' + selectCargo.Amount + '<br>' +
                        'CargoDetail: ' + selectCargo.CargoDetail;
        cargoData.appendChild(div);
    }
}