/****************************************
 Javascript middleware for sharespot
 Version: 20211015 0250
 Authored by Yuwei Le
 ****************************************/

const host = 'http://119.3.189.237:8240/';
const route = 'api/v1/';

function get(api, callback) {
    $.ajax({
        url: host + route + api,
        dataType: 'text',
        cache: false,
        contentType: false,
        processData: false,
        type: 'get',
        success: function (data) {
            var ret = JSON.parse(data);
            if (ret.recode !== '0') {
                callback(1, '获取成功但发送错误');
                return;
            }
            callback(0, ret.data);
        },
        error: function(){
            callback(1, '发生错误');
        }
    });
}

function post(api, postdata, callback) {
    $.ajax({
        url: host + route + api,
        dataType: 'text',
        cache: false,
        contentType: false,
        processData: false,
        data: postdata,
        type: 'post',
        success: function (data) {
            var ret = JSON.parse(data);
            if (ret.recode !== '0') {
                callback(1, '获取成功但发送错误');
                return;
            }
            callback(0, ret.data);
        },
        error: function(){
            callback(1, '发生错误');
        }
    });
}

function trim(s){
    return trimRight(trimLeft(s));
}

function trimLeft(s){
    if(s == null) {
        return "";
    }
    var whitespace = new String(" \t\n\r");
    var str = new String(s);
    if (whitespace.indexOf(str.charAt(0)) != -1) {
        var j=0, i = str.length;
        while (j < i && whitespace.indexOf(str.charAt(j)) != -1){
            j++;
        }
        str = str.substring(j, i);
    }
    return str;
}

function trimRight(s){
    if(s == null) return "";
    var whitespace = new String(" \t\n\r");
    var str = new String(s);
    if (whitespace.indexOf(str.charAt(str.length-1)) != -1){
        var i = str.length - 1;
        while (i >= 0 && whitespace.indexOf(str.charAt(i)) != -1){
            i--;
        }
        str = str.substring(0, i+1);
    }
    return str;
}