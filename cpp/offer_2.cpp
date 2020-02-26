//
// Created by polar on 2020/2/26.
//

#include <cstring>
#include <iostream>
#include "offer_2.h"

CMyString & CMyString::operator=(CMyString &str){
    if (this == &str) {
        return *this;
    }

    delete []m_pData;
    m_pData = nullptr;

    m_pData = new char[strlen(str.m_pData) + 1];
    strcpy(m_pData, str.m_pData);

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