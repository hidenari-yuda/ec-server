    document.getElementsByClassName("confirmDelete").addEventListener('click', function() {
        const result = confirm('削除しますか');
        if(result){
          alert('削除しました');
  
        }else{
          alert('削除をとりやめました');
        }
      })