document.addEventListener('DOMContentLoaded', function() {

var categoryFirst = document.getElementsByClassName("category_first");
var text = categoryFirst[0].innerText
var categoryStr = categoryString(text)
categoryFirst[0].innerText = categoryStr;

var categorySecond = document.getElementsByClassName("category_second");
var text = categorySecond[0].innerText
var categoryStr = categoryString(text)
categorySecond[0].innerText = categoryStr;

var categoryThird = document.getElementsByClassName("category_third");
var text = categoryThird[0].innerText
var categoryStr = categoryString(text)
categoryThird[0].innerText = categoryStr;

function categoryString(category){
  if(category == "0"){
    return "テント・タープ";
  }else if (category == "1"){
    return "チェア・テーブル";
  }else if(category == "2"){
    return "セット用品";
  }else if(category == "3"){
    return "調理用品";
  }else if(category == "4"){
    return "寝袋・マット";
  }else if(category == "5"){
    return "ランタン";
  }else if(category == "6"){
    return "クーラーボックス";
  }else if(category == "7"){
    return "焚き火";
  }else if(category == "8"){
    return "その他";
  }else{
    return "";
  }
}

})