//cat /sys/devices/ocp.2/helper.14/AIN6
// echo 

// I got PWM working on my BBB from the command line the other day.  Here's what I did on Angstrom (4/13 image).  

// Add the am33xx_pwm module to the bone_capemgr.  You'll have to look at your BBB for the # of the cape manager.  Mine is bone_capemgr.8, but since I've only tested on a single board, I'm not sure if the numbers are the same across every board.
// $ echo am33xx_pwm > /sys/devices/bone_capemgr.#/slots

// Add the individual PWM module to the bone_capemgr. For example, EHRPWM2B is pin 13 on the P8 connector.  To enable this pin for PWM, use the following command
// $ echo bone_pwm_P8_13 > /sys/devices/bone_capemgr.#/slots

// This creates a link in /sys/devices/ocp.#/pwm_test_P8_13.#.  For my BBB, I have /sys/devices/ocp.2/pwm_test_P8_13.12.  The number in the pwm_test has changed for me for a couple times that I've used it, so you'll have to see what number it assigns for you.

// All the settings for the PWM configuration are found in that folder similar to the old interface.  To generate a 50 Hz signal with 50% duty cycle, (the period and duty times are giving in nanoseconds)
// $ echo 20000000 > /sys/devices/ocp.#/pwm_test_P8_13.#/period 
// $ echo 10000000 > /sys/devices/ocp.#/pwm_test_P8_13.#/duty

// I found all this from digging through the BoneScript source code found here https://github.com/jadonk/bonescript.  bonescript/node_modules/bonescript/index.js has the core implementation of the module.

// Hope this helps.