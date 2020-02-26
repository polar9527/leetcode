//
// Created by polar on 2020/2/26.
//

#ifndef LEETCODE_OFFER_2_H
#define LEETCODE_OFFER_2_H


class CMyString {
friend std::ostream & operator<<(std::ostream &os, const CMyString &str);
public:
    explicit CMyString(char *pData = nullptr){
        if (pData == nullptr)
            return ;
        m_pData = new char[strlen(pData) + 1];
        strcpy(m_pData, pData);
    };
    CMyString(const CMyString &str){
        m_pData = new char[strlen(str.m_pData) + 1];
        strcpy(m_pData, str.m_pData);
    };
    CMyString &operator=(CMyString &str);
    ~CMyString(){};

private:
    char *m_pData;
};



#endif //LEETCODE_OFFER_2_H
