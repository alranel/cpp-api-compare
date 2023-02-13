# API comparison

1. [Global Functions](#global-functions)
2. [Client](#client)
3. [IPAddress](#ipaddress)
4. [Print](#print)
5. [Server](#server)
6. [Stream](#stream)
7. [String](#string)
8. [UDP](#udp)
9. [Serial](#serial)
10. [SPI](#spi)
11. [Wire](#wire)

## Global Functions

|                 |                                                                                 | AVR        | ESP32  |
| --------------- | ------------------------------------------------------------------------------- | ---------- | ------ |
|                 | **Fully compatible**                                                            |            |        |
| `void`          | `analogWrite(uint8_t pin, int value)`                                           | ✔️         | ✔️     |
| `void`          | `analogWrite(uint8_t pin, int val)`                                             | ✔️         | ✔️     |
| `void`          | `attachInterrupt(uint8_t pin, std::function< void(void)> intRoutine, int mode)` | ✔️         | ✔️     |
| `void`          | `attachInterrupt(uint8_t pin, void(*)(void) handler, int mode)`                 | ✔️         | ✔️     |
| `void`          | `attachInterrupt(uint8_t interruptNum, void(*)(void) userFunc, int mode)`       | ✔️         | ✔️     |
| `void`          | `delay(unsigned long ms)`                                                       | ✔️         | ✔️     |
| `void`          | `delay(uint32_t ms)`                                                            | ✔️         | ✔️     |
| `void`          | `delayMicroseconds(unsigned int us)`                                            | ✔️         | ✔️     |
| `void`          | `delayMicroseconds(uint32_t us)`                                                | ✔️         | ✔️     |
| `void`          | `detachInterrupt(uint8_t pin)`                                                  | ✔️         | ✔️     |
| `void`          | `detachInterrupt(uint8_t interruptNum)`                                         | ✔️         | ✔️     |
| `int`           | `digitalRead(uint8_t pin)`                                                      | ✔️         | ✔️     |
| `void`          | `digitalWrite(uint8_t pin, uint8_t val)`                                        | ✔️         | ✔️     |
| `boolean`       | `isAlpha(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isAlphaNumeric(int c)`                                                         | ✔️         | ✔️     |
| `boolean`       | `isAscii(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isControl(int c)`                                                              | ✔️         | ✔️     |
| `boolean`       | `isDigit(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isGraph(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isHexadecimalDigit(int c)`                                                     | ✔️         | ✔️     |
| `boolean`       | `isLowerCase(int c)`                                                            | ✔️         | ✔️     |
| `boolean`       | `isPrintable(int c)`                                                            | ✔️         | ✔️     |
| `boolean`       | `isPunct(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isSpace(int c)`                                                                | ✔️         | ✔️     |
| `boolean`       | `isUpperCase(int c)`                                                            | ✔️         | ✔️     |
| `boolean`       | `isWhitespace(int c)`                                                           | ✔️         | ✔️     |
| `long`          | `map(long x, long in_min, long in_max, long out_min, long out_max)`             | ✔️         | ✔️     |
| `unsigned long` | `micros()`                                                                      | ✔️         | ✔️     |
| `unsigned long` | `millis()`                                                                      | ✔️         | ✔️     |
| `void`          | `noTone(uint8_t _pin)`                                                          | ✔️         | ✔️     |
| `void`          | `pinMode(uint8_t pin, uint8_t mode)`                                            | ✔️         | ✔️     |
| `unsigned long` | `pulseIn(uint8_t pin, uint8_t state, unsigned long timeout)`                    | ✔️         | ✔️     |
| `unsigned long` | `pulseInLong(uint8_t pin, uint8_t state, unsigned long timeout)`                | ✔️         | ✔️     |
| `long`          | `random(long howbig)`                                                           | ✔️         | ✔️     |
| `long`          | `random(long howsmall, long howbig)`                                            | ✔️         | ✔️     |
| `void`          | `randomSeed(unsigned long seed)`                                                | ✔️         | ✔️     |
| `uint8_t`       | `shiftIn(uint8_t dataPin, uint8_t clockPin, uint8_t bitOrder)`                  | ✔️         | ✔️     |
| `void`          | `shiftOut(uint8_t dataPin, uint8_t clockPin, uint8_t bitOrder, uint8_t val)`    | ✔️         | ✔️     |
| `void`          | `tone(uint8_t _pin, unsigned int frequency, unsigned long duration = 0)`        | ✔️         | ✔️     |
|                 | **Differing return values**                                                     |            |        |
|                 | `analogRead(uint8_t pin)`                                                       | `uint16_t` | `int`  |
|                 | **Limited compatibility**                                                       |            |        |
| `void`          | `analogReadResolution(uint8_t bits)`                                            |            | ✔️     |
| `void`          | `analogReference(uint8_t mode)`                                                 | ✔️         |        |

## Client

|                 |                                                                                               | AVR    | ESP32                |
| --------------- | --------------------------------------------------------------------------------------------- | ------ | -------------------- |
|                 | **Fully compatible**                                                                          |        |                      |
| `int`           | `available()`                                                                                 | ✔️     | ✔️                   |
| `void`          | `clearWriteError()`                                                                           | ✔️     | ✔️                   |
| `int`           | `connect(const char * host, uint16_t port)`                                                   | ✔️     | ✔️                   |
| `int`           | `connect(IPAddress ip, uint16_t port)`                                                        | ✔️     | ✔️                   |
| `uint8_t`       | `connected()`                                                                                 | ✔️     | ✔️                   |
| `bool`          | `find(char target)`                                                                           | ✔️     | ✔️                   |
| `bool`          | `find(const char * target)`                                                                   | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target)`                                                                      | ✔️     | ✔️                   |
| `bool`          | `find(char * target)`                                                                         | ✔️     | ✔️                   |
| `bool`          | `find(const char * target, size_t length)`                                                    | ✔️     | ✔️                   |
| `bool`          | `find(const uint8_t * target, size_t length)`                                                 | ✔️     | ✔️                   |
| `bool`          | `find(char * target, size_t length)`                                                          | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target, size_t length)`                                                       | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, const char * terminator)`                                     | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, char * terminator)`                                              | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, const char * terminator)`                                  | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, char * terminator)`                                                 | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, size_t targetLen, char * terminate, size_t termLen)`                | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, size_t targetLen, const char * terminate, size_t termLen)`    | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, size_t targetLen, const char * terminate, size_t termLen)` | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, size_t targetLen, char * terminate, size_t termLen)`             | ✔️     | ✔️                   |
| `void`          | `flush()`                                                                                     | ✔️     | ✔️                   |
| `unsigned long` | `getTimeout()`                                                                                | ✔️     | ✔️                   |
| `int`           | `getWriteError()`                                                                             | ✔️     | ✔️                   |
|                 | `operator bool()`                                                                             | ✔️     | ✔️                   |
| `int`           | `peek()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(char )`                                                                                | ✔️     | ✔️                   |
| `size_t`        | `print(const char )`                                                                          | ✔️     | ✔️                   |
| `size_t`        | `print(const  &Printable )`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `print(const  &String )`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(const __FlashStringHelper * )`                                                         | ✔️     | ✔️                   |
| `size_t`        | `print(int , int  = 10)`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(double , int  = 2)`                                                                    | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned char , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long long , int  = 10)`                                                       | ✔️     | ✔️                   |
| `size_t`        | `print(struct tm * timeinfo, const char * format = NULL)`                                     | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned int , int  = 10)`                                                             | ✔️     | ✔️                   |
| `size_t`        | `print(long , int  = 10)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `print(long long , int  = 10)`                                                                | ✔️     | ✔️                   |
| `size_t`        | `println()`                                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const  &Printable )`                                                                 | ✔️     | ✔️                   |
| `size_t`        | `println(char )`                                                                              | ✔️     | ✔️                   |
| `size_t`        | `println(const  &String s)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const char )`                                                                        | ✔️     | ✔️                   |
| `size_t`        | `println(const __FlashStringHelper * )`                                                       | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned char , int  = 10)`                                                          | ✔️     | ✔️                   |
| `size_t`        | `println(long , int  = 10)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned int , int  = 10)`                                                           | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long long , int  = 10)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `println(double , int  = 2)`                                                                  | ✔️     | ✔️                   |
| `size_t`        | `println(long long , int  = 10)`                                                              | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long , int  = 10)`                                                          | ✔️     | ✔️                   |
| `size_t`        | `println(struct tm * timeinfo, const char * format = NULL)`                                   | ✔️     | ✔️                   |
| `size_t`        | `println(int , int  = 10)`                                                                    | ✔️     | ✔️                   |
| `int`           | `read()`                                                                                      | ✔️     | ✔️                   |
| `int`           | `read(uint8_t * buf, size_t size)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `readBytes(char * buffer, size_t length)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `readBytes(uint8_t * buffer, size_t length)`                                                  | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, char * buffer, size_t length)`                               | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, uint8_t * buffer, size_t length)`                            | ✔️     | ✔️                   |
| `String`        | `readString()`                                                                                | ✔️     | ✔️                   |
| `String`        | `readStringUntil(char terminator)`                                                            | ✔️     | ✔️                   |
| `void`          | `setTimeout(unsigned long timeout)`                                                           | ✔️     | ✔️                   |
| `void`          | `stop()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `write(uint8_t )`                                                                             | ✔️     | ✔️                   |
| `size_t`        | `write(const char * str)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(const char * buffer, size_t size)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(const uint8_t * buf, size_t size)`                                                     | ✔️     | ✔️                   |
|                 | **Differing return values**                                                                   |        |                      |
|                 | `availableForWrite()`                                                                         | `int`  | `size_t virtual int` |
|                 | **Limited compatibility**                                                                     |        |                      |
| `float`         | `parseFloat()`                                                                                |        | ✔️                   |
| `float`         | `parseFloat(LookaheadMode lookahead, char ignore = '\x01')`                                   | ✔️     |                      |
| `long`          | `parseInt()`                                                                                  |        | ✔️                   |
| `long`          | `parseInt(LookaheadMode lookahead, char ignore = '\x01')`                                     | ✔️     |                      |
| `size_t`        | `printf(const char * format, ... )`                                                           |        | ✔️                   |

## IPAddress

|               |                                                                                                   | AVR       | ESP32       |
| ------------- | ------------------------------------------------------------------------------------------------- | --------- | ----------- |
|               | **Fully compatible**                                                                              |           |             |
|               | `IPAddress()`                                                                                     | ✔️        | ✔️          |
|               | `IPAddress(uint32_t address)`                                                                     | ✔️        | ✔️          |
|               | `IPAddress(const uint8_t * address)`                                                              | ✔️        | ✔️          |
|               | `IPAddress(uint8_t first_octet, uint8_t second_octet, uint8_t third_octet, uint8_t fourth_octet)` | ✔️        | ✔️          |
| `bool`        | `fromString(const  &String address)`                                                              | ✔️        | ✔️          |
| `bool`        | `fromString(const char * address)`                                                                | ✔️        | ✔️          |
|               | `operator uint32_t()`                                                                             | ✔️        | ✔️          |
| ` &IPAddress` | `operator=(uint32_t address)`                                                                     | ✔️        | ✔️          |
| ` &IPAddress` | `operator=(const uint8_t * address)`                                                              | ✔️        | ✔️          |
| `bool`        | `operator==(const  &IPAddress addr)`                                                              | ✔️        | ✔️          |
| `bool`        | `operator==(const uint8_t * addr)`                                                                | ✔️        | ✔️          |
| `size_t`      | `printTo( &Print p)`                                                                              | ✔️        | ✔️          |
|               | **Differing return values**                                                                       |           |             |
|               | `operator[](int index)`                                                                           | `uint8_t` | `uint8_t &` |
|               | **Limited compatibility**                                                                         |           |             |
| `String`      | `toString()`                                                                                      |           | ✔️          |

## Print

|          |                                                             | AVR    | ESP32                |
| -------- | ----------------------------------------------------------- | ------ | -------------------- |
|          | **Fully compatible**                                        |        |                      |
|          | `Print()`                                                   | ✔️     | ✔️                   |
| `void`   | `clearWriteError()`                                         | ✔️     | ✔️                   |
| `void`   | `flush()`                                                   | ✔️     | ✔️                   |
| `int`    | `getWriteError()`                                           | ✔️     | ✔️                   |
| `size_t` | `print(const  &Printable )`                                 | ✔️     | ✔️                   |
| `size_t` | `print(const __FlashStringHelper * )`                       | ✔️     | ✔️                   |
| `size_t` | `print(const  &String )`                                    | ✔️     | ✔️                   |
| `size_t` | `print(const char )`                                        | ✔️     | ✔️                   |
| `size_t` | `print(char )`                                              | ✔️     | ✔️                   |
| `size_t` | `print(struct tm * timeinfo, const char * format = NULL)`   | ✔️     | ✔️                   |
| `size_t` | `print(unsigned long long , int  = 10)`                     | ✔️     | ✔️                   |
| `size_t` | `print(int , int  = 10)`                                    | ✔️     | ✔️                   |
| `size_t` | `print(long , int  = 10)`                                   | ✔️     | ✔️                   |
| `size_t` | `print(double , int  = 2)`                                  | ✔️     | ✔️                   |
| `size_t` | `print(unsigned char , int  = 10)`                          | ✔️     | ✔️                   |
| `size_t` | `print(long long , int  = 10)`                              | ✔️     | ✔️                   |
| `size_t` | `print(unsigned long , int  = 10)`                          | ✔️     | ✔️                   |
| `size_t` | `print(unsigned int , int  = 10)`                           | ✔️     | ✔️                   |
| `size_t` | `println()`                                                 | ✔️     | ✔️                   |
| `size_t` | `println(const  &Printable )`                               | ✔️     | ✔️                   |
| `size_t` | `println(const  &String s)`                                 | ✔️     | ✔️                   |
| `size_t` | `println(const char )`                                      | ✔️     | ✔️                   |
| `size_t` | `println(const __FlashStringHelper * )`                     | ✔️     | ✔️                   |
| `size_t` | `println(char )`                                            | ✔️     | ✔️                   |
| `size_t` | `println(unsigned long , int  = 10)`                        | ✔️     | ✔️                   |
| `size_t` | `println(struct tm * timeinfo, const char * format = NULL)` | ✔️     | ✔️                   |
| `size_t` | `println(long long , int  = 10)`                            | ✔️     | ✔️                   |
| `size_t` | `println(unsigned long long , int  = 10)`                   | ✔️     | ✔️                   |
| `size_t` | `println(long , int  = 10)`                                 | ✔️     | ✔️                   |
| `size_t` | `println(unsigned char , int  = 10)`                        | ✔️     | ✔️                   |
| `size_t` | `println(int , int  = 10)`                                  | ✔️     | ✔️                   |
| `size_t` | `println(unsigned int , int  = 10)`                         | ✔️     | ✔️                   |
| `size_t` | `println(double , int  = 2)`                                | ✔️     | ✔️                   |
| `size_t` | `write(const char * str)`                                   | ✔️     | ✔️                   |
| `size_t` | `write(uint8_t )`                                           | ✔️     | ✔️                   |
| `size_t` | `write(const char * buffer, size_t size)`                   | ✔️     | ✔️                   |
| `size_t` | `write(const uint8_t * buffer, size_t size)`                | ✔️     | ✔️                   |
|          | **Differing return values**                                 |        |                      |
|          | `availableForWrite()`                                       | `int`  | `size_t virtual int` |
|          | **Limited compatibility**                                   |        |                      |
| `size_t` | `printf(const char * format, ... )`                         |        | ✔️                   |

## Server

|          |                                                             | AVR                  | ESP32  |
| -------- | ----------------------------------------------------------- | -------------------- | ------ |
|          | **Fully compatible**                                        |                      |        |
| `void`   | `begin()`                                                   | ✔️                   | ✔️     |
| `void`   | `clearWriteError()`                                         | ✔️                   | ✔️     |
| `void`   | `flush()`                                                   | ✔️                   | ✔️     |
| `int`    | `getWriteError()`                                           | ✔️                   | ✔️     |
| `size_t` | `print(const char )`                                        | ✔️                   | ✔️     |
| `size_t` | `print(const __FlashStringHelper * )`                       | ✔️                   | ✔️     |
| `size_t` | `print(const  &Printable )`                                 | ✔️                   | ✔️     |
| `size_t` | `print(char )`                                              | ✔️                   | ✔️     |
| `size_t` | `print(const  &String )`                                    | ✔️                   | ✔️     |
| `size_t` | `print(struct tm * timeinfo, const char * format = NULL)`   | ✔️                   | ✔️     |
| `size_t` | `print(long long , int  = 10)`                              | ✔️                   | ✔️     |
| `size_t` | `print(int , int  = 10)`                                    | ✔️                   | ✔️     |
| `size_t` | `print(unsigned char , int  = 10)`                          | ✔️                   | ✔️     |
| `size_t` | `print(unsigned long , int  = 10)`                          | ✔️                   | ✔️     |
| `size_t` | `print(unsigned long long , int  = 10)`                     | ✔️                   | ✔️     |
| `size_t` | `print(unsigned int , int  = 10)`                           | ✔️                   | ✔️     |
| `size_t` | `print(double , int  = 2)`                                  | ✔️                   | ✔️     |
| `size_t` | `print(long , int  = 10)`                                   | ✔️                   | ✔️     |
| `size_t` | `println()`                                                 | ✔️                   | ✔️     |
| `size_t` | `println(const  &Printable )`                               | ✔️                   | ✔️     |
| `size_t` | `println(const char )`                                      | ✔️                   | ✔️     |
| `size_t` | `println(const __FlashStringHelper * )`                     | ✔️                   | ✔️     |
| `size_t` | `println(char )`                                            | ✔️                   | ✔️     |
| `size_t` | `println(const  &String s)`                                 | ✔️                   | ✔️     |
| `size_t` | `println(unsigned int , int  = 10)`                         | ✔️                   | ✔️     |
| `size_t` | `println(int , int  = 10)`                                  | ✔️                   | ✔️     |
| `size_t` | `println(long long , int  = 10)`                            | ✔️                   | ✔️     |
| `size_t` | `println(double , int  = 2)`                                | ✔️                   | ✔️     |
| `size_t` | `println(unsigned char , int  = 10)`                        | ✔️                   | ✔️     |
| `size_t` | `println(unsigned long , int  = 10)`                        | ✔️                   | ✔️     |
| `size_t` | `println(struct tm * timeinfo, const char * format = NULL)` | ✔️                   | ✔️     |
| `size_t` | `println(unsigned long long , int  = 10)`                   | ✔️                   | ✔️     |
| `size_t` | `println(long , int  = 10)`                                 | ✔️                   | ✔️     |
| `size_t` | `write(const char * str)`                                   | ✔️                   | ✔️     |
| `size_t` | `write(uint8_t )`                                           | ✔️                   | ✔️     |
| `size_t` | `write(const uint8_t * buffer, size_t size)`                | ✔️                   | ✔️     |
| `size_t` | `write(const char * buffer, size_t size)`                   | ✔️                   | ✔️     |
|          | **Differing return values**                                 |                      |        |
|          | `availableForWrite()`                                       | `size_t virtual int` | `int`  |
|          | **Limited compatibility**                                   |                      |        |
| `void`   | `begin(uint16_t port = 0)`                                  |                      | ✔️     |
| `size_t` | `printf(const char * format, ... )`                         |                      | ✔️     |

## Stream

|                 |                                                                                               | AVR    | ESP32                |
| --------------- | --------------------------------------------------------------------------------------------- | ------ | -------------------- |
|                 | **Fully compatible**                                                                          |        |                      |
|                 | `Stream()`                                                                                    | ✔️     | ✔️                   |
| `int`           | `available()`                                                                                 | ✔️     | ✔️                   |
| `void`          | `clearWriteError()`                                                                           | ✔️     | ✔️                   |
| `bool`          | `find(char target)`                                                                           | ✔️     | ✔️                   |
| `bool`          | `find(char * target)`                                                                         | ✔️     | ✔️                   |
| `bool`          | `find(const char * target)`                                                                   | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target)`                                                                      | ✔️     | ✔️                   |
| `bool`          | `find(const uint8_t * target, size_t length)`                                                 | ✔️     | ✔️                   |
| `bool`          | `find(const char * target, size_t length)`                                                    | ✔️     | ✔️                   |
| `bool`          | `find(char * target, size_t length)`                                                          | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target, size_t length)`                                                       | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, char * terminator)`                                                 | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, const char * terminator)`                                  | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, char * terminator)`                                              | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, const char * terminator)`                                     | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, size_t targetLen, char * terminate, size_t termLen)`                | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, size_t targetLen, const char * terminate, size_t termLen)` | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, size_t targetLen, char * terminate, size_t termLen)`             | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, size_t targetLen, const char * terminate, size_t termLen)`    | ✔️     | ✔️                   |
| `void`          | `flush()`                                                                                     | ✔️     | ✔️                   |
| `unsigned long` | `getTimeout()`                                                                                | ✔️     | ✔️                   |
| `int`           | `getWriteError()`                                                                             | ✔️     | ✔️                   |
| `int`           | `peek()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(const char )`                                                                          | ✔️     | ✔️                   |
| `size_t`        | `print(const __FlashStringHelper * )`                                                         | ✔️     | ✔️                   |
| `size_t`        | `print(const  &Printable )`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `print(const  &String )`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(char )`                                                                                | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long long , int  = 10)`                                                       | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(long long , int  = 10)`                                                                | ✔️     | ✔️                   |
| `size_t`        | `print(double , int  = 2)`                                                                    | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned int , int  = 10)`                                                             | ✔️     | ✔️                   |
| `size_t`        | `print(long , int  = 10)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned char , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(int , int  = 10)`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(struct tm * timeinfo, const char * format = NULL)`                                     | ✔️     | ✔️                   |
| `size_t`        | `println()`                                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const  &String s)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const  &Printable )`                                                                 | ✔️     | ✔️                   |
| `size_t`        | `println(const char )`                                                                        | ✔️     | ✔️                   |
| `size_t`        | `println(char )`                                                                              | ✔️     | ✔️                   |
| `size_t`        | `println(const __FlashStringHelper * )`                                                       | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned int , int  = 10)`                                                           | ✔️     | ✔️                   |
| `size_t`        | `println(double , int  = 2)`                                                                  | ✔️     | ✔️                   |
| `size_t`        | `println(long , int  = 10)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long long , int  = 10)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned char , int  = 10)`                                                          | ✔️     | ✔️                   |
| `size_t`        | `println(int , int  = 10)`                                                                    | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long , int  = 10)`                                                          | ✔️     | ✔️                   |
| `size_t`        | `println(struct tm * timeinfo, const char * format = NULL)`                                   | ✔️     | ✔️                   |
| `size_t`        | `println(long long , int  = 10)`                                                              | ✔️     | ✔️                   |
| `int`           | `read()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `readBytes(uint8_t * buffer, size_t length)`                                                  | ✔️     | ✔️                   |
| `size_t`        | `readBytes(char * buffer, size_t length)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, char * buffer, size_t length)`                               | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, uint8_t * buffer, size_t length)`                            | ✔️     | ✔️                   |
| `String`        | `readString()`                                                                                | ✔️     | ✔️                   |
| `String`        | `readStringUntil(char terminator)`                                                            | ✔️     | ✔️                   |
| `void`          | `setTimeout(unsigned long timeout)`                                                           | ✔️     | ✔️                   |
| `size_t`        | `write(uint8_t )`                                                                             | ✔️     | ✔️                   |
| `size_t`        | `write(const char * str)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(const char * buffer, size_t size)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(const uint8_t * buffer, size_t size)`                                                  | ✔️     | ✔️                   |
|                 | **Differing return values**                                                                   |        |                      |
|                 | `availableForWrite()`                                                                         | `int`  | `size_t virtual int` |
|                 | **Limited compatibility**                                                                     |        |                      |
| `float`         | `parseFloat()`                                                                                |        | ✔️                   |
| `float`         | `parseFloat(LookaheadMode lookahead, char ignore = '\x01')`                                   | ✔️     |                      |
| `long`          | `parseInt()`                                                                                  |        | ✔️                   |
| `long`          | `parseInt(LookaheadMode lookahead, char ignore = '\x01')`                                     |        | ✔️                   |
| `size_t`        | `printf(const char * format, ... )`                                                           | ✔️     |                      |

## String

|                 |                                                                                  | AVR      | ESP32          |
| --------------- | -------------------------------------------------------------------------------- | -------- | -------------- |
|                 | **Fully compatible**                                                             |          |                |
|                 | `String(const char * cstr = "")`                                                 | ✔️       | ✔️             |
|                 | `String(char c)`                                                                 | ✔️       | ✔️             |
|                 | `String(const __FlashStringHelper * str)`                                        | ✔️       | ✔️             |
|                 | `String(const  &String str)`                                                     | ✔️       | ✔️             |
|                 | `String(long , unsigned char base = 10)`                                         | ✔️       | ✔️             |
|                 | `String(unsigned long , unsigned char base = 10)`                                | ✔️       | ✔️             |
|                 | `String(double , unsigned char decimalPlaces = 2)`                               | ✔️       | ✔️             |
|                 | `String(const char * cstr, unsigned int length)`                                 | ✔️       | ✔️             |
|                 | `String(double , unsigned int decimalPlaces = 2)`                                | ✔️       | ✔️             |
|                 | `String(unsigned int , unsigned char base = 10)`                                 | ✔️       | ✔️             |
|                 | `String(float , unsigned int decimalPlaces = 2)`                                 | ✔️       | ✔️             |
|                 | `String(unsigned char , unsigned char base = 10)`                                | ✔️       | ✔️             |
|                 | `String(int , unsigned char base = 10)`                                          | ✔️       | ✔️             |
|                 | `String(unsigned long long , unsigned char base = 10)`                           | ✔️       | ✔️             |
|                 | `String(float , unsigned char decimalPlaces = 2)`                                | ✔️       | ✔️             |
|                 | `String(long long , unsigned char base = 10)`                                    | ✔️       | ✔️             |
| `const char *`  | `c_str()`                                                                        | ✔️       | ✔️             |
| `char`          | `charAt(unsigned int index)`                                                     | ✔️       | ✔️             |
| `int`           | `compareTo(const  &String s)`                                                    | ✔️       | ✔️             |
| `unsigned char` | `concat(const char * cstr)`                                                      | ✔️       | ✔️             |
| `unsigned char` | `concat(unsigned char c)`                                                        | ✔️       | ✔️             |
| `unsigned char` | `concat(unsigned long long num)`                                                 | ✔️       | ✔️             |
| `unsigned char` | `concat(float num)`                                                              | ✔️       | ✔️             |
| `unsigned char` | `concat(const __FlashStringHelper * str)`                                        | ✔️       | ✔️             |
| `unsigned char` | `concat(double num)`                                                             | ✔️       | ✔️             |
| `unsigned char` | `concat(int num)`                                                                | ✔️       | ✔️             |
| `unsigned char` | `concat(char c)`                                                                 | ✔️       | ✔️             |
| `unsigned char` | `concat(const  &String str)`                                                     | ✔️       | ✔️             |
| `unsigned char` | `concat(long long num)`                                                          | ✔️       | ✔️             |
| `unsigned char` | `concat(unsigned long num)`                                                      | ✔️       | ✔️             |
| `unsigned char` | `concat(long num)`                                                               | ✔️       | ✔️             |
| `unsigned char` | `concat(unsigned int num)`                                                       | ✔️       | ✔️             |
| `unsigned char` | `endsWith(const char * suffix)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `endsWith(const  &String suffix)`                                                | ✔️       | ✔️             |
| `unsigned char` | `endsWith(const __FlashStringHelper * suffix)`                                   | ✔️       | ✔️             |
| `unsigned char` | `equals(const  &String s)`                                                       | ✔️       | ✔️             |
| `unsigned char` | `equals(const char * cstr)`                                                      | ✔️       | ✔️             |
| `unsigned char` | `equalsIgnoreCase(const  &String s)`                                             | ✔️       | ✔️             |
| `void`          | `getBytes(unsigned char * buf, unsigned int bufsize, unsigned int index = 0)`    | ✔️       | ✔️             |
| `int`           | `indexOf(char ch)`                                                               | ✔️       | ✔️             |
| `int`           | `indexOf(const  &String str)`                                                    | ✔️       | ✔️             |
| `int`           | `indexOf(const  &String str, unsigned int fromIndex)`                            | ✔️       | ✔️             |
| `int`           | `indexOf(char ch, unsigned int fromIndex)`                                       | ✔️       | ✔️             |
| `int`           | `lastIndexOf(const  &String str)`                                                | ✔️       | ✔️             |
| `int`           | `lastIndexOf(char ch)`                                                           | ✔️       | ✔️             |
| `int`           | `lastIndexOf(char ch, unsigned int fromIndex)`                                   | ✔️       | ✔️             |
| `int`           | `lastIndexOf(const  &String str, unsigned int fromIndex)`                        | ✔️       | ✔️             |
| `unsigned int`  | `length()`                                                                       | ✔️       | ✔️             |
|                 | `operator StringIfHelperType()`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator!=(const char * cstr)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator!=(const  &String rhs)`                                                 | ✔️       | ✔️             |
| ` &String`      | `operator+=(float num)`                                                          | ✔️       | ✔️             |
| ` &String`      | `operator+=(unsigned int num)`                                                   | ✔️       | ✔️             |
| ` &String`      | `operator+=(int num)`                                                            | ✔️       | ✔️             |
| ` &String`      | `operator+=(unsigned long num)`                                                  | ✔️       | ✔️             |
| ` &String`      | `operator+=(double num)`                                                         | ✔️       | ✔️             |
| ` &String`      | `operator+=(unsigned long long num)`                                             | ✔️       | ✔️             |
| ` &String`      | `operator+=(char c)`                                                             | ✔️       | ✔️             |
| ` &String`      | `operator+=(long num)`                                                           | ✔️       | ✔️             |
| ` &String`      | `operator+=(long long num)`                                                      | ✔️       | ✔️             |
| ` &String`      | `operator+=(const __FlashStringHelper * str)`                                    | ✔️       | ✔️             |
| ` &String`      | `operator+=(const  &String rhs)`                                                 | ✔️       | ✔️             |
| ` &String`      | `operator+=(const char * cstr)`                                                  | ✔️       | ✔️             |
| ` &String`      | `operator+=(unsigned char num)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator<(const  &String rhs)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator<=(const  &String rhs)`                                                 | ✔️       | ✔️             |
| ` &String`      | `operator=(const  &String rhs)`                                                  | ✔️       | ✔️             |
| ` &String`      | `operator=(const char * cstr)`                                                   | ✔️       | ✔️             |
| ` &String`      | `operator=(const __FlashStringHelper * str)`                                     | ✔️       | ✔️             |
| `unsigned char` | `operator==(const  &String rhs)`                                                 | ✔️       | ✔️             |
| `unsigned char` | `operator==(const char * cstr)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator>(const  &String rhs)`                                                  | ✔️       | ✔️             |
| `unsigned char` | `operator>=(const  &String rhs)`                                                 | ✔️       | ✔️             |
| `void`          | `remove(unsigned int index)`                                                     | ✔️       | ✔️             |
| `void`          | `remove(unsigned int index, unsigned int count)`                                 | ✔️       | ✔️             |
| `void`          | `replace(const __FlashStringHelper * find, const char * replace)`                | ✔️       | ✔️             |
| `void`          | `replace(char find, char replace)`                                               | ✔️       | ✔️             |
| `void`          | `replace(const  &String find, const  &String replace)`                           | ✔️       | ✔️             |
| `void`          | `replace(const __FlashStringHelper * find, const __FlashStringHelper * replace)` | ✔️       | ✔️             |
| `void`          | `replace(const char * find, const char * replace)`                               | ✔️       | ✔️             |
| `void`          | `replace(const __FlashStringHelper * find, const  &String replace)`              | ✔️       | ✔️             |
| `void`          | `replace(const char * find, const  &String replace)`                             | ✔️       | ✔️             |
| `unsigned char` | `reserve(unsigned int size)`                                                     | ✔️       | ✔️             |
| `void`          | `setCharAt(unsigned int index, char c)`                                          | ✔️       | ✔️             |
| `unsigned char` | `startsWith(const __FlashStringHelper * prefix)`                                 | ✔️       | ✔️             |
| `unsigned char` | `startsWith(const  &String prefix)`                                              | ✔️       | ✔️             |
| `unsigned char` | `startsWith(const char * prefix)`                                                | ✔️       | ✔️             |
| `unsigned char` | `startsWith(const  &String prefix, unsigned int offset)`                         | ✔️       | ✔️             |
| `String`        | `substring(unsigned int beginIndex)`                                             | ✔️       | ✔️             |
| `String`        | `substring(unsigned int beginIndex, unsigned int endIndex)`                      | ✔️       | ✔️             |
| `void`          | `toCharArray(char * buf, unsigned int bufsize, unsigned int index = 0)`          | ✔️       | ✔️             |
| `double`        | `toDouble()`                                                                     | ✔️       | ✔️             |
| `float`         | `toFloat()`                                                                      | ✔️       | ✔️             |
| `long`          | `toInt()`                                                                        | ✔️       | ✔️             |
| `void`          | `toLowerCase()`                                                                  | ✔️       | ✔️             |
| `void`          | `toUpperCase()`                                                                  | ✔️       | ✔️             |
| `void`          | `trim()`                                                                         | ✔️       | ✔️             |
|                 | **Differing return values**                                                      |          |                |
|                 | `begin()`                                                                        | `char *` | `const char *` |
|                 | `end()`                                                                          | `char *` | `const char *` |
|                 | `operator[](unsigned int index)`                                                 | `char`   | `char &`       |
|                 | **Limited compatibility**                                                        |          |                |
| `void`          | `clear()`                                                                        |          | ✔️             |
| `unsigned char` | `concat(const char * cstr, unsigned int length)`                                 |          | ✔️             |
| `unsigned char` | `concat(const uint8_t * cstr, unsigned int length)`                              |          | ✔️             |
| `unsigned char` | `equalsConstantTime(const  &String s)`                                           |          | ✔️             |
| `bool`          | `isEmpty()`                                                                      |          | ✔️             |

## UDP

|                 |                                                                                               | AVR    | ESP32                |
| --------------- | --------------------------------------------------------------------------------------------- | ------ | -------------------- |
|                 | **Fully compatible**                                                                          |        |                      |
| `int`           | `available()`                                                                                 | ✔️     | ✔️                   |
| `uint8_t`       | `begin(uint16_t )`                                                                            | ✔️     | ✔️                   |
| `uint8_t`       | `beginMulticast(IPAddress , uint16_t )`                                                       | ✔️     | ✔️                   |
| `int`           | `beginPacket(IPAddress ip, uint16_t port)`                                                    | ✔️     | ✔️                   |
| `int`           | `beginPacket(const char * host, uint16_t port)`                                               | ✔️     | ✔️                   |
| `void`          | `clearWriteError()`                                                                           | ✔️     | ✔️                   |
| `int`           | `endPacket()`                                                                                 | ✔️     | ✔️                   |
| `bool`          | `find(char target)`                                                                           | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target)`                                                                      | ✔️     | ✔️                   |
| `bool`          | `find(char * target)`                                                                         | ✔️     | ✔️                   |
| `bool`          | `find(const char * target)`                                                                   | ✔️     | ✔️                   |
| `bool`          | `find(char * target, size_t length)`                                                          | ✔️     | ✔️                   |
| `bool`          | `find(const uint8_t * target, size_t length)`                                                 | ✔️     | ✔️                   |
| `bool`          | `find(const char * target, size_t length)`                                                    | ✔️     | ✔️                   |
| `bool`          | `find(uint8_t * target, size_t length)`                                                       | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, const char * terminator)`                                     | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, char * terminator)`                                                 | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, const char * terminator)`                                  | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, char * terminator)`                                              | ✔️     | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, size_t targetLen, char * terminate, size_t termLen)`             | ✔️     | ✔️                   |
| `bool`          | `findUntil(const char * target, size_t targetLen, const char * terminate, size_t termLen)`    | ✔️     | ✔️                   |
| `bool`          | `findUntil(char * target, size_t targetLen, char * terminate, size_t termLen)`                | ✔️     | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, size_t targetLen, const char * terminate, size_t termLen)` | ✔️     | ✔️                   |
| `void`          | `flush()`                                                                                     | ✔️     | ✔️                   |
| `unsigned long` | `getTimeout()`                                                                                | ✔️     | ✔️                   |
| `int`           | `getWriteError()`                                                                             | ✔️     | ✔️                   |
| `int`           | `parsePacket()`                                                                               | ✔️     | ✔️                   |
| `int`           | `peek()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(const __FlashStringHelper * )`                                                         | ✔️     | ✔️                   |
| `size_t`        | `print(char )`                                                                                | ✔️     | ✔️                   |
| `size_t`        | `print(const  &Printable )`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `print(const  &String )`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(const char )`                                                                          | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned char , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(int , int  = 10)`                                                                      | ✔️     | ✔️                   |
| `size_t`        | `print(struct tm * timeinfo, const char * format = NULL)`                                     | ✔️     | ✔️                   |
| `size_t`        | `print(double , int  = 2)`                                                                    | ✔️     | ✔️                   |
| `size_t`        | `print(long , int  = 10)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long , int  = 10)`                                                            | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned int , int  = 10)`                                                             | ✔️     | ✔️                   |
| `size_t`        | `print(unsigned long long , int  = 10)`                                                       | ✔️     | ✔️                   |
| `size_t`        | `print(long long , int  = 10)`                                                                | ✔️     | ✔️                   |
| `size_t`        | `println()`                                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const __FlashStringHelper * )`                                                       | ✔️     | ✔️                   |
| `size_t`        | `println(const  &String s)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(const  &Printable )`                                                                 | ✔️     | ✔️                   |
| `size_t`        | `println(const char )`                                                                        | ✔️     | ✔️                   |
| `size_t`        | `println(char )`                                                                              | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long long , int  = 10)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `println(long long , int  = 10)`                                                              | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned int , int  = 10)`                                                           | ✔️     | ✔️                   |
| `size_t`        | `println(int , int  = 10)`                                                                    | ✔️     | ✔️                   |
| `size_t`        | `println(struct tm * timeinfo, const char * format = NULL)`                                   | ✔️     | ✔️                   |
| `size_t`        | `println(long , int  = 10)`                                                                   | ✔️     | ✔️                   |
| `size_t`        | `println(double , int  = 2)`                                                                  | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned char , int  = 10)`                                                          | ✔️     | ✔️                   |
| `size_t`        | `println(unsigned long , int  = 10)`                                                          | ✔️     | ✔️                   |
| `int`           | `read()`                                                                                      | ✔️     | ✔️                   |
| `int`           | `read(char * buffer, size_t len)`                                                             | ✔️     | ✔️                   |
| `int`           | `read(unsigned char * buffer, size_t len)`                                                    | ✔️     | ✔️                   |
| `size_t`        | `readBytes(char * buffer, size_t length)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `readBytes(uint8_t * buffer, size_t length)`                                                  | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, char * buffer, size_t length)`                               | ✔️     | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, uint8_t * buffer, size_t length)`                            | ✔️     | ✔️                   |
| `String`        | `readString()`                                                                                | ✔️     | ✔️                   |
| `String`        | `readStringUntil(char terminator)`                                                            | ✔️     | ✔️                   |
| `IPAddress`     | `remoteIP()`                                                                                  | ✔️     | ✔️                   |
| `uint16_t`      | `remotePort()`                                                                                | ✔️     | ✔️                   |
| `void`          | `setTimeout(unsigned long timeout)`                                                           | ✔️     | ✔️                   |
| `void`          | `stop()`                                                                                      | ✔️     | ✔️                   |
| `size_t`        | `write(const char * str)`                                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(uint8_t )`                                                                             | ✔️     | ✔️                   |
| `size_t`        | `write(const char * buffer, size_t size)`                                                     | ✔️     | ✔️                   |
| `size_t`        | `write(const uint8_t * buffer, size_t size)`                                                  | ✔️     | ✔️                   |
|                 | **Differing return values**                                                                   |        |                      |
|                 | `availableForWrite()`                                                                         | `int`  | `size_t virtual int` |
|                 | **Limited compatibility**                                                                     |        |                      |
| `float`         | `parseFloat()`                                                                                |        | ✔️                   |
| `float`         | `parseFloat(LookaheadMode lookahead, char ignore = '\x01')`                                   | ✔️     |                      |
| `long`          | `parseInt()`                                                                                  |        | ✔️                   |
| `long`          | `parseInt(LookaheadMode lookahead, char ignore = '\x01')`                                     | ✔️     |                      |
| `size_t`        | `printf(const char * format, ... )`                                                           |        | ✔️                   |

## Serial

|                 |                                                                                                                                                                                          | AVR    | ESP32  |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ------ |
|                 | **Fully compatible**                                                                                                                                                                     |        |        |
| `int`           | `available()`                                                                                                                                                                            | ✔️     | ✔️     |
| `int`           | `availableForWrite()`                                                                                                                                                                    | ✔️     | ✔️     |
| `void`          | `begin(unsigned long baud)`                                                                                                                                                              | ✔️     | ✔️     |
| `void`          | `begin(unsigned long , uint8_t )`                                                                                                                                                        | ✔️     | ✔️     |
| `void`          | `clearWriteError()`                                                                                                                                                                      | ✔️     | ✔️     |
| `void`          | `end()`                                                                                                                                                                                  | ✔️     | ✔️     |
| `bool`          | `find(char target)`                                                                                                                                                                      | ✔️     | ✔️     |
| `bool`          | `find(char * target)`                                                                                                                                                                    | ✔️     | ✔️     |
| `bool`          | `find(const char * target)`                                                                                                                                                              | ✔️     | ✔️     |
| `bool`          | `find(uint8_t * target)`                                                                                                                                                                 | ✔️     | ✔️     |
| `bool`          | `find(const char * target, size_t length)`                                                                                                                                               | ✔️     | ✔️     |
| `bool`          | `find(const uint8_t * target, size_t length)`                                                                                                                                            | ✔️     | ✔️     |
| `bool`          | `find(char * target, size_t length)`                                                                                                                                                     | ✔️     | ✔️     |
| `bool`          | `find(uint8_t * target, size_t length)`                                                                                                                                                  | ✔️     | ✔️     |
| `bool`          | `findUntil(uint8_t * target, char * terminator)`                                                                                                                                         | ✔️     | ✔️     |
| `bool`          | `findUntil(char * target, char * terminator)`                                                                                                                                            | ✔️     | ✔️     |
| `bool`          | `findUntil(const uint8_t * target, const char * terminator)`                                                                                                                             | ✔️     | ✔️     |
| `bool`          | `findUntil(const char * target, const char * terminator)`                                                                                                                                | ✔️     | ✔️     |
| `bool`          | `findUntil(const char * target, size_t targetLen, const char * terminate, size_t termLen)`                                                                                               | ✔️     | ✔️     |
| `bool`          | `findUntil(char * target, size_t targetLen, char * terminate, size_t termLen)`                                                                                                           | ✔️     | ✔️     |
| `bool`          | `findUntil(const uint8_t * target, size_t targetLen, const char * terminate, size_t termLen)`                                                                                            | ✔️     | ✔️     |
| `bool`          | `findUntil(uint8_t * target, size_t targetLen, char * terminate, size_t termLen)`                                                                                                        | ✔️     | ✔️     |
| `void`          | `flush()`                                                                                                                                                                                | ✔️     | ✔️     |
| `unsigned long` | `getTimeout()`                                                                                                                                                                           | ✔️     | ✔️     |
| `int`           | `getWriteError()`                                                                                                                                                                        | ✔️     | ✔️     |
|                 | `operator bool()`                                                                                                                                                                        | ✔️     | ✔️     |
| `int`           | `peek()`                                                                                                                                                                                 | ✔️     | ✔️     |
| `size_t`        | `print(const __FlashStringHelper * )`                                                                                                                                                    | ✔️     | ✔️     |
| `size_t`        | `print(const  &String )`                                                                                                                                                                 | ✔️     | ✔️     |
| `size_t`        | `print(const char )`                                                                                                                                                                     | ✔️     | ✔️     |
| `size_t`        | `print(const  &Printable )`                                                                                                                                                              | ✔️     | ✔️     |
| `size_t`        | `print(char )`                                                                                                                                                                           | ✔️     | ✔️     |
| `size_t`        | `print(long , int  = 10)`                                                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `print(unsigned char , int  = 10)`                                                                                                                                                       | ✔️     | ✔️     |
| `size_t`        | `print(double , int  = 2)`                                                                                                                                                               | ✔️     | ✔️     |
| `size_t`        | `print(struct tm * timeinfo, const char * format = NULL)`                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `print(unsigned int , int  = 10)`                                                                                                                                                        | ✔️     | ✔️     |
| `size_t`        | `print(int , int  = 10)`                                                                                                                                                                 | ✔️     | ✔️     |
| `size_t`        | `print(long long , int  = 10)`                                                                                                                                                           | ✔️     | ✔️     |
| `size_t`        | `print(unsigned long long , int  = 10)`                                                                                                                                                  | ✔️     | ✔️     |
| `size_t`        | `print(unsigned long , int  = 10)`                                                                                                                                                       | ✔️     | ✔️     |
| `size_t`        | `println()`                                                                                                                                                                              | ✔️     | ✔️     |
| `size_t`        | `println(char )`                                                                                                                                                                         | ✔️     | ✔️     |
| `size_t`        | `println(const  &String s)`                                                                                                                                                              | ✔️     | ✔️     |
| `size_t`        | `println(const __FlashStringHelper * )`                                                                                                                                                  | ✔️     | ✔️     |
| `size_t`        | `println(const char )`                                                                                                                                                                   | ✔️     | ✔️     |
| `size_t`        | `println(const  &Printable )`                                                                                                                                                            | ✔️     | ✔️     |
| `size_t`        | `println(long , int  = 10)`                                                                                                                                                              | ✔️     | ✔️     |
| `size_t`        | `println(unsigned char , int  = 10)`                                                                                                                                                     | ✔️     | ✔️     |
| `size_t`        | `println(double , int  = 2)`                                                                                                                                                             | ✔️     | ✔️     |
| `size_t`        | `println(long long , int  = 10)`                                                                                                                                                         | ✔️     | ✔️     |
| `size_t`        | `println(unsigned long , int  = 10)`                                                                                                                                                     | ✔️     | ✔️     |
| `size_t`        | `println(struct tm * timeinfo, const char * format = NULL)`                                                                                                                              | ✔️     | ✔️     |
| `size_t`        | `println(unsigned long long , int  = 10)`                                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `println(unsigned int , int  = 10)`                                                                                                                                                      | ✔️     | ✔️     |
| `size_t`        | `println(int , int  = 10)`                                                                                                                                                               | ✔️     | ✔️     |
| `int`           | `read()`                                                                                                                                                                                 | ✔️     | ✔️     |
| `size_t`        | `readBytes(uint8_t * buffer, size_t length)`                                                                                                                                             | ✔️     | ✔️     |
| `size_t`        | `readBytes(char * buffer, size_t length)`                                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `readBytesUntil(char terminator, uint8_t * buffer, size_t length)`                                                                                                                       | ✔️     | ✔️     |
| `size_t`        | `readBytesUntil(char terminator, char * buffer, size_t length)`                                                                                                                          | ✔️     | ✔️     |
| `String`        | `readString()`                                                                                                                                                                           | ✔️     | ✔️     |
| `String`        | `readStringUntil(char terminator)`                                                                                                                                                       | ✔️     | ✔️     |
| `void`          | `setTimeout(unsigned long timeout)`                                                                                                                                                      | ✔️     | ✔️     |
| `size_t`        | `write(unsigned int n)`                                                                                                                                                                  | ✔️     | ✔️     |
| `size_t`        | `write(unsigned long n)`                                                                                                                                                                 | ✔️     | ✔️     |
| `size_t`        | `write(long n)`                                                                                                                                                                          | ✔️     | ✔️     |
| `size_t`        | `write(int n)`                                                                                                                                                                           | ✔️     | ✔️     |
| `size_t`        | `write(const char * s)`                                                                                                                                                                  | ✔️     | ✔️     |
| `size_t`        | `write(uint8_t )`                                                                                                                                                                        | ✔️     | ✔️     |
| `size_t`        | `write(const char * str)`                                                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `write(const char * buffer, size_t size)`                                                                                                                                                | ✔️     | ✔️     |
| `size_t`        | `write(const uint8_t * buffer, size_t size)`                                                                                                                                             | ✔️     | ✔️     |
|                 | **Limited compatibility**                                                                                                                                                                |        |        |
| `void`          | `_rx_complete_irq()`                                                                                                                                                                     | ✔️     |        |
| `void`          | `_tx_udr_empty_irq()`                                                                                                                                                                    | ✔️     |        |
| `uint32_t`      | `baudRate()`                                                                                                                                                                             |        | ✔️     |
| `void`          | `begin(unsigned long baud, uint32_t config = 0x800001c, int8_t rxPin = -1, int8_t txPin = -1, bool invert = false, unsigned long timeout_ms = 20000UL, uint8_t rxfifo_full_thrhd = 112)` |        | ✔️     |
| `void`          | `end(bool fullyTerminate = true)`                                                                                                                                                        | ✔️     |        |
| `void`          | `eventQueueReset()`                                                                                                                                                                      |        | ✔️     |
| `void`          | `flush(bool txOnly)`                                                                                                                                                                     |        | ✔️     |
| `void`          | `onReceive(OnReceiveCb function, bool onlyOnTimeout = false)`                                                                                                                            |        | ✔️     |
| `void`          | `onReceiveError(OnReceiveErrorCb function)`                                                                                                                                              |        | ✔️     |
| `float`         | `parseFloat()`                                                                                                                                                                           |        | ✔️     |
| `float`         | `parseFloat(LookaheadMode lookahead, char ignore = '\x01')`                                                                                                                              | ✔️     |        |
| `long`          | `parseInt()`                                                                                                                                                                             |        | ✔️     |
| `long`          | `parseInt(LookaheadMode lookahead, char ignore = '\x01')`                                                                                                                                |        | ✔️     |
| `size_t`        | `printf(const char * format, ... )`                                                                                                                                                      |        | ✔️     |
| `size_t`        | `read(char * buffer, size_t size)`                                                                                                                                                       |        | ✔️     |
| `size_t`        | `read(uint8_t * buffer, size_t size)`                                                                                                                                                    |        | ✔️     |
| `void`          | `setDebugOutput(bool )`                                                                                                                                                                  |        | ✔️     |
| `void`          | `setHwFlowCtrlMode(uint8_t mode = 0x3, uint8_t threshold = 64)`                                                                                                                          | ✔️     |        |
| `void`          | `setPins(int8_t rxPin, int8_t txPin, int8_t ctsPin = -1, int8_t rtsPin = -1)`                                                                                                            |        | ✔️     |
| `size_t`        | `setRxBufferSize(size_t new_size)`                                                                                                                                                       |        | ✔️     |
| `void`          | `setRxFIFOFull(uint8_t fifoBytes)`                                                                                                                                                       |        | ✔️     |
| `void`          | `setRxInvert(bool )`                                                                                                                                                                     |        | ✔️     |
| `void`          | `setRxTimeout(uint8_t symbols_timeout)`                                                                                                                                                  |        | ✔️     |
| `size_t`        | `setTxBufferSize(size_t new_size)`                                                                                                                                                       |        | ✔️     |
| `void`          | `updateBaudRate(unsigned long baud)`                                                                                                                                                     | ✔️     |        |

## SPI

|            |                                                                              | AVR    | ESP32  |
| ---------- | ---------------------------------------------------------------------------- | ------ | ------ |
|            | **Fully compatible**                                                         |        |        |
| `void`     | `begin()`                                                                    | ✔️     | ✔️     |
| `void`     | `beginTransaction(SPISettings settings)`                                     | ✔️     | ✔️     |
| `void`     | `end()`                                                                      | ✔️     | ✔️     |
| `void`     | `endTransaction()`                                                           | ✔️     | ✔️     |
| `void`     | `setBitOrder(uint8_t bitOrder)`                                              | ✔️     | ✔️     |
| `void`     | `setClockDivider(uint8_t clockDiv)`                                          | ✔️     | ✔️     |
| `void`     | `setClockDivider(uint32_t clockDiv)`                                         | ✔️     | ✔️     |
| `void`     | `setDataMode(uint8_t dataMode)`                                              | ✔️     | ✔️     |
| `uint8_t`  | `transfer(uint8_t data)`                                                     | ✔️     | ✔️     |
| `void`     | `transfer(void * data, uint32_t size)`                                       | ✔️     | ✔️     |
| `void`     | `transfer(void * buf, size_t count)`                                         | ✔️     | ✔️     |
| `uint16_t` | `transfer16(uint16_t data)`                                                  | ✔️     | ✔️     |
|            | **Limited compatibility**                                                    |        |        |
| `void`     | `attachInterrupt()`                                                          | ✔️     |        |
| `void`     | `begin(int8_t sck = -1, int8_t miso = -1, int8_t mosi = -1, int8_t ss = -1)` |        | ✔️     |
| ` *spi_t`  | `bus()`                                                                      | ✔️     |        |
| `void`     | `detachInterrupt()`                                                          | ✔️     |        |
| `uint32_t` | `getClockDivider()`                                                          |        | ✔️     |
| `void`     | `notUsingInterrupt(uint8_t interruptNumber)`                                 | ✔️     |        |
| `int8_t`   | `pinSS()`                                                                    |        | ✔️     |
| `void`     | `setFrequency(uint32_t freq)`                                                |        | ✔️     |
| `void`     | `setHwCs(bool use)`                                                          |        | ✔️     |
| `uint32_t` | `transfer32(uint32_t data)`                                                  |        | ✔️     |
| `void`     | `transferBits(uint32_t data, uint32_t * out, uint8_t bits)`                  | ✔️     |        |
| `void`     | `transferBytes(const uint8_t * data, uint8_t * out, uint32_t size)`          |        | ✔️     |
| `void`     | `usingInterrupt(uint8_t interruptNumber)`                                    |        | ✔️     |
| `void`     | `write(uint8_t data)`                                                        |        | ✔️     |
| `void`     | `write16(uint16_t data)`                                                     |        | ✔️     |
| `void`     | `write32(uint32_t data)`                                                     |        | ✔️     |
| `void`     | `writeBytes(const uint8_t * data, uint32_t size)`                            |        | ✔️     |
| `void`     | `writePattern(const uint8_t * data, uint8_t size, uint32_t repeat)`          | ✔️     |        |
| `void`     | `writePixels(const void * data, uint32_t size)`                              |        | ✔️     |

## Wire

|                 |                                                                                               | AVR       | ESP32                |
| --------------- | --------------------------------------------------------------------------------------------- | --------- | -------------------- |
|                 | **Fully compatible**                                                                          |           |                      |
| `int`           | `available()`                                                                                 | ✔️        | ✔️                   |
| `void`          | `beginTransmission(uint8_t )`                                                                 | ✔️        | ✔️                   |
| `void`          | `beginTransmission(uint16_t address)`                                                         | ✔️        | ✔️                   |
| `void`          | `beginTransmission(int address)`                                                              | ✔️        | ✔️                   |
| `void`          | `beginTransmission(int )`                                                                     | ✔️        | ✔️                   |
| `void`          | `beginTransmission(uint8_t address)`                                                          | ✔️        | ✔️                   |
| `void`          | `clearWriteError()`                                                                           | ✔️        | ✔️                   |
| `uint8_t`       | `endTransmission()`                                                                           | ✔️        | ✔️                   |
| `uint8_t`       | `endTransmission(bool sendStop)`                                                              | ✔️        | ✔️                   |
| `uint8_t`       | `endTransmission(uint8_t )`                                                                   | ✔️        | ✔️                   |
| `bool`          | `find(char target)`                                                                           | ✔️        | ✔️                   |
| `bool`          | `find(uint8_t * target)`                                                                      | ✔️        | ✔️                   |
| `bool`          | `find(char * target)`                                                                         | ✔️        | ✔️                   |
| `bool`          | `find(const char * target)`                                                                   | ✔️        | ✔️                   |
| `bool`          | `find(const uint8_t * target, size_t length)`                                                 | ✔️        | ✔️                   |
| `bool`          | `find(char * target, size_t length)`                                                          | ✔️        | ✔️                   |
| `bool`          | `find(const char * target, size_t length)`                                                    | ✔️        | ✔️                   |
| `bool`          | `find(uint8_t * target, size_t length)`                                                       | ✔️        | ✔️                   |
| `bool`          | `findUntil(char * target, char * terminator)`                                                 | ✔️        | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, char * terminator)`                                              | ✔️        | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, const char * terminator)`                                  | ✔️        | ✔️                   |
| `bool`          | `findUntil(const char * target, const char * terminator)`                                     | ✔️        | ✔️                   |
| `bool`          | `findUntil(const uint8_t * target, size_t targetLen, const char * terminate, size_t termLen)` | ✔️        | ✔️                   |
| `bool`          | `findUntil(uint8_t * target, size_t targetLen, char * terminate, size_t termLen)`             | ✔️        | ✔️                   |
| `bool`          | `findUntil(const char * target, size_t targetLen, const char * terminate, size_t termLen)`    | ✔️        | ✔️                   |
| `bool`          | `findUntil(char * target, size_t targetLen, char * terminate, size_t termLen)`                | ✔️        | ✔️                   |
| `void`          | `flush()`                                                                                     | ✔️        | ✔️                   |
| `unsigned long` | `getTimeout()`                                                                                | ✔️        | ✔️                   |
| `int`           | `getWriteError()`                                                                             | ✔️        | ✔️                   |
| `void`          | `onReceive(void(*)(int) )`                                                                    | ✔️        | ✔️                   |
| `void`          | `onRequest(void(*)(void) )`                                                                   | ✔️        | ✔️                   |
| `int`           | `peek()`                                                                                      | ✔️        | ✔️                   |
| `size_t`        | `print(const __FlashStringHelper * )`                                                         | ✔️        | ✔️                   |
| `size_t`        | `print(char )`                                                                                | ✔️        | ✔️                   |
| `size_t`        | `print(const char )`                                                                          | ✔️        | ✔️                   |
| `size_t`        | `print(const  &Printable )`                                                                   | ✔️        | ✔️                   |
| `size_t`        | `print(const  &String )`                                                                      | ✔️        | ✔️                   |
| `size_t`        | `print(long , int  = 10)`                                                                     | ✔️        | ✔️                   |
| `size_t`        | `print(int , int  = 10)`                                                                      | ✔️        | ✔️                   |
| `size_t`        | `print(unsigned long long , int  = 10)`                                                       | ✔️        | ✔️                   |
| `size_t`        | `print(long long , int  = 10)`                                                                | ✔️        | ✔️                   |
| `size_t`        | `print(unsigned int , int  = 10)`                                                             | ✔️        | ✔️                   |
| `size_t`        | `print(unsigned long , int  = 10)`                                                            | ✔️        | ✔️                   |
| `size_t`        | `print(struct tm * timeinfo, const char * format = NULL)`                                     | ✔️        | ✔️                   |
| `size_t`        | `print(unsigned char , int  = 10)`                                                            | ✔️        | ✔️                   |
| `size_t`        | `print(double , int  = 2)`                                                                    | ✔️        | ✔️                   |
| `size_t`        | `println()`                                                                                   | ✔️        | ✔️                   |
| `size_t`        | `println(const char )`                                                                        | ✔️        | ✔️                   |
| `size_t`        | `println(char )`                                                                              | ✔️        | ✔️                   |
| `size_t`        | `println(const  &Printable )`                                                                 | ✔️        | ✔️                   |
| `size_t`        | `println(const  &String s)`                                                                   | ✔️        | ✔️                   |
| `size_t`        | `println(const __FlashStringHelper * )`                                                       | ✔️        | ✔️                   |
| `size_t`        | `println(unsigned char , int  = 10)`                                                          | ✔️        | ✔️                   |
| `size_t`        | `println(int , int  = 10)`                                                                    | ✔️        | ✔️                   |
| `size_t`        | `println(long long , int  = 10)`                                                              | ✔️        | ✔️                   |
| `size_t`        | `println(unsigned long , int  = 10)`                                                          | ✔️        | ✔️                   |
| `size_t`        | `println(double , int  = 2)`                                                                  | ✔️        | ✔️                   |
| `size_t`        | `println(struct tm * timeinfo, const char * format = NULL)`                                   | ✔️        | ✔️                   |
| `size_t`        | `println(unsigned int , int  = 10)`                                                           | ✔️        | ✔️                   |
| `size_t`        | `println(long , int  = 10)`                                                                   | ✔️        | ✔️                   |
| `size_t`        | `println(unsigned long long , int  = 10)`                                                     | ✔️        | ✔️                   |
| `int`           | `read()`                                                                                      | ✔️        | ✔️                   |
| `size_t`        | `readBytes(char * buffer, size_t length)`                                                     | ✔️        | ✔️                   |
| `size_t`        | `readBytes(uint8_t * buffer, size_t length)`                                                  | ✔️        | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, char * buffer, size_t length)`                               | ✔️        | ✔️                   |
| `size_t`        | `readBytesUntil(char terminator, uint8_t * buffer, size_t length)`                            | ✔️        | ✔️                   |
| `String`        | `readString()`                                                                                | ✔️        | ✔️                   |
| `String`        | `readStringUntil(char terminator)`                                                            | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint16_t address, uint8_t size)`                                                 | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(int address, int size)`                                                          | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint8_t , uint8_t )`                                                             | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(int , int )`                                                                     | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint8_t address, uint8_t size)`                                                  | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint8_t address, uint8_t size, uint8_t sendStop)`                                | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint16_t address, uint8_t size, uint8_t sendStop)`                               | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(uint16_t address, uint8_t size, bool sendStop)`                                  | ✔️        | ✔️                   |
| `uint8_t`       | `requestFrom(int address, int size, int sendStop)`                                            | ✔️        | ✔️                   |
| `void`          | `setTimeout(unsigned long timeout)`                                                           | ✔️        | ✔️                   |
| `size_t`        | `write(int n)`                                                                                | ✔️        | ✔️                   |
| `size_t`        | `write(const char * s)`                                                                       | ✔️        | ✔️                   |
| `size_t`        | `write(unsigned long n)`                                                                      | ✔️        | ✔️                   |
| `size_t`        | `write(unsigned int n)`                                                                       | ✔️        | ✔️                   |
| `size_t`        | `write(uint8_t )`                                                                             | ✔️        | ✔️                   |
| `size_t`        | `write(long n)`                                                                               | ✔️        | ✔️                   |
| `size_t`        | `write(const char * str)`                                                                     | ✔️        | ✔️                   |
| `size_t`        | `write(const uint8_t * buffer, size_t size)`                                                  | ✔️        | ✔️                   |
| `size_t`        | `write(const uint8_t * , size_t )`                                                            | ✔️        | ✔️                   |
| `size_t`        | `write(const char * buffer, size_t size)`                                                     | ✔️        | ✔️                   |
|                 | **Differing return values**                                                                   |           |                      |
|                 | `availableForWrite()`                                                                         | `int`     | `size_t virtual int` |
|                 | `begin()`                                                                                     | `void`    | `bool`               |
|                 | `begin(int addr)`                                                                             | `void`    | `bool`               |
|                 | `begin(int )`                                                                                 | `void`    | `bool`               |
|                 | `begin(uint8_t addr)`                                                                         | `void`    | `bool`               |
|                 | `begin(uint8_t )`                                                                             | `void`    | `bool`               |
|                 | `end()`                                                                                       | `void`    | `bool`               |
|                 | `requestFrom(uint8_t , uint8_t , uint8_t )`                                                   | `uint8_t` | `size_t`             |
|                 | `requestFrom(uint8_t address, size_t len, bool stopBit)`                                      | `uint8_t` | `size_t`             |
|                 | `requestFrom(uint16_t address, size_t size, bool sendStop)`                                   | `uint8_t` | `size_t`             |
|                 | `requestFrom(int , int , int )`                                                               | `uint8_t` | `size_t`             |
|                 | `setClock(uint32_t )`                                                                         | `void`    | `bool`               |
|                 | **Limited compatibility**                                                                     |           |                      |
| `bool`          | `begin(int sda, int scl, uint32_t frequency = 0)`                                             |           | ✔️                   |
| `bool`          | `begin(uint8_t slaveAddr, int sda, int scl, uint32_t frequency)`                              |           | ✔️                   |
| `void`          | `clearWireTimeoutFlag()`                                                                      | ✔️        |                      |
| `uint32_t`      | `getClock()`                                                                                  |           | ✔️                   |
| `uint16_t`      | `getTimeOut()`                                                                                |           | ✔️                   |
| `bool`          | `getWireTimeoutFlag()`                                                                        | ✔️        |                      |
| `float`         | `parseFloat()`                                                                                |           | ✔️                   |
| `float`         | `parseFloat(LookaheadMode lookahead, char ignore = '\x01')`                                   | ✔️        |                      |
| `long`          | `parseInt()`                                                                                  |           | ✔️                   |
| `long`          | `parseInt(LookaheadMode lookahead, char ignore = '\x01')`                                     | ✔️        |                      |
| `size_t`        | `printf(const char * format, ... )`                                                           | ✔️        |                      |
| `uint8_t`       | `requestFrom(uint8_t , uint8_t , uint32_t , uint8_t , uint8_t )`                              | ✔️        |                      |
| `size_t`        | `setBufferSize(size_t bSize)`                                                                 |           | ✔️                   |
| `bool`          | `setPins(int sda, int scl)`                                                                   |           | ✔️                   |
| `void`          | `setTimeOut(uint16_t timeOutMillis)`                                                          |           | ✔️                   |
| `void`          | `setWireTimeout(uint32_t timeout = 25000, bool reset_with_timeout = false)`                   | ✔️        |                      |
| `size_t`        | `slaveWrite(const uint8_t * , size_t )`                                                       |           | ✔️                   |

