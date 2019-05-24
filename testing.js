

function aw(){
    var array = [];
    for (var i = 0; i< 60; i ++){
        array.push('The actors are here! ', + i);
     
     }

     for (var i = 0; i< 60; i ++){
        console.log(array[i])
     
     }
}
// However, the cue is not announced until at least 5000ms have
// passed through the use of setTimeout
setTimeout(function() {  
    return aw()
}, 3000);

// This console log is executed right away
console.log('An exploration of art and music. And now, as we wait for the actors...');  