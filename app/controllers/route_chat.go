package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hidenari-yuda/todo_app/app/models"
)

//チャットのグループの処理
/*func chatGroup(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		chatgroups, err := user.GetChatGroupsByUser()
		if err != nil {
			log.Println(err)
		}
		user.ChatGroups = chatgroups
		fmt.Println(user.ChatGroups)
		generateHTML(w, user, "layout", "private_navbar", "chat")
	}
}

func chatGroupNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "chatGroup_new")
	}

}

func chatGroupSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		chat_owner, chat_member, chat_name := r.PostFormValue("chat_owner"), r.PostFormValue("chat_member"), r.PostFormValue("chat_name")
		if err := user.CreateChatGroup(chat_owner, chat_member, chat_name); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}
}
func chatGroupEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetChat(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "chat_edit")

	}
}

func chatGroupUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Chat{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateChatGroup(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}
}

func chatGroupDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err = sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		cg, err := models.GetChatGroup(id)
		if err != nil {
			log.Println(err)
		}
		if err := cg.DeleteChatGroup(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}

}*/

//チャットのコンテンツの処理

func chat(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		chats, err := user.GetChatsByUser()
		if err != nil {
			log.Println(err)
		}
		user.Chats = chats
		fmt.Println(user.Chats)
		generateHTML(w, user, "layout", "private_navbar", "chat")
	}
}

func chatNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "chat_new")
	}

}

func chatSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateChat(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}
}
func chatEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetChat(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "chat_edit")

	}
}

func chatUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Chat{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateChat(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}
}

func chatDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err = sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetChat(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteChat(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/chat", 302)
	}

}
