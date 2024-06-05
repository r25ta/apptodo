function removeFromDb(id){
    fetch(`/delete?id=${id}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
 }
 
 function updateDb(id, item) {
    let input = document.getElementById(item)
    let newitem = input.value
    fetch(`/update?id=${id}&olditem=${item}&newitem=${newitem}`, {method: "PUT"}).then(res =>{
        if (res.status == 200){
        alert("Database updated")
            window.location.pathname = "/"
        }
    })
 }