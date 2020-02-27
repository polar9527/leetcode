//
// Created by polar on 2020/2/26.
//

#include <cstring>
#include <iostream>
#include "offer_2.h"

CMyString & CMyString::operator=(CMyString &str){
    if (this != &str) {
        CMyString strTemp(str);

        char *pTemp = strTemp.m_pData;
        strTemp.m_pData = m_pData;
        m_pData = pTemp;
    }

    return *this;
}

std::ostream & operator<<(std::ostream &os, const CMyString &str) {
    os << str.m_pData;
    return os;
}

int main(){
    char c1[] = "c1";
    CMyString s1(c1);
    CMyString s2 = s1;
    std::cout << s2 << std::endl;

}