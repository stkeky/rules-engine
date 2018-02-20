import tornado.web


class Handler(tornado.web.RequestHandler):
    def get(self, user_id):
        self.render("editor.html", user_id=user_id)
