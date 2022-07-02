export function cookieRead(doc) {
  var cookiesAll = doc.cookie;
  var cookieArray = cookiesAll.split(';');
  var cookies = new Map();

  for(var i = 0; i < cookieArray.length; i++) {
    var name = cookieArray[i].split('=')[0].trim();
    var value = cookieArray[i].split('=')[1];

    cookies.set(name, value);
  }

  console.log("cookies");
  console.log(cookies);

  return cookies;
}

export function cookieDeleteSesssion(doc) {
  var now = new Date();
  now.setMonth(now.getMonth() - 1);

  doc.cookie = "session=";
  doc.cookie = "expires=" + now.toUTCString() + ";"
}
