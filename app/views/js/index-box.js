function body() {
    var dates = document.getElementsByName("created");

    var len = dates.length;

    for (var i = 0; i < len; i++){
        datesText = dates[i].innerText;

        var subDates = datesText.substring(0, 11);

        dates[i].innerText = subDates;
        }
}

// createdAtを表示形式変更
/*function dateFormat() {
    const elements = document.getElementsByClassName(`todo-created`);
    Array.prototype.filter.call(elements, function (el) {
      const resDate = new Date(el.innerText)
      const year = resDate.getFullYear()
      const month = resDate.getMonth() + 1
      const day = resDate.getDate()
      const dateString = `${year}/${month}/${day}`
      el.textContent = dateString
    })
  }*/