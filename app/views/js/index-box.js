function body() {
    var dates = document.getElementsByName("created");

    var len = dates.length;

    for (var i = 0; i < len; i++){
        datesText = dates[i].innerText;

        var subDates = datesText.substring(0, 11);

        dates[i].innerText = subDates;
        }
}