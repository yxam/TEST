ReactDOM.render(React.createElement(
  "div",
  { "class": "container" },
  React.createElement(
    "form",
    { "class": "form-signin" },
    React.createElement(
      "h2",
      { "class": "form-sigin-heading" },
      "SigIn"
    ),
    React.createElement(
      "label",
      { "class": "sr-only" },
      "Rut"
    ),
    React.createElement("input", { type: "text", "class": "form-control", name: "rut", placeholder: "18.***.***-*" }),
    React.createElement(
      "label",
      { "class": "sr-only" },
      "Password"
    ),
    React.createElement("input", { type: "text", "class": "form-control", name: "pass", placeholder: "********" }),
    React.createElement(
      "button",
      { "class": "btn btn-primary", type: "submit", name: "login" },
      "Sigin"
    )
  )
), document.getElementById('example'));