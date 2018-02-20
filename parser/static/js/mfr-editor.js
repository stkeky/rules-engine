function Editor(elementName) {
    this.element = document.getElementById(elementName);
    this.editor = {};
    this.init = function () {
        CodeMirror.defineSimpleMode("mfrules", {
            start: [
                {
                    regex: /"(?:[^\\]|\\.)*?(?:"|$)/,
                    token: "string"
                },
                {
                    regex: /(?:SEND EMAIL|TURN OFF|TO)\b/,
                    token: "action"
                },
                {
                    regex: /([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})/,
                    token: "uuid"
                },
                {
                    regex: /[-+\/*=<>!\[\]]+/,
                    token: "operator"
                },
                {
                    regex: /(?:BETWEEN)\b/,
                    token: "operator2"
                },
                {
                    regex: /(?:RULE|TRIGGERS)\b/,
                    token: "keyword",
                    indent: true
                },
                {
                    regex: /true|false|null/,
                    token: "atom"
                },
                {
                    regex: /0x[a-f\d]+|[-+]?(?:\.\d+|\d+\.?\d*)(?:e[-+]?\d+)?/i,
                    token: "number"
                },
                {
                    regex: /[a-z$][\w$]*/,
                    token: "variable"
                }
            ]
        });

        this.editor = CodeMirror.fromTextArea(this.element, {
            lineNumbers: true,
            mode: "mfrules",
            theme: "xq-light"
        });

        this.editor.setSize(null, 600);
    }
}

