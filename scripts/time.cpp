#include <iostream>
#include <iomanip>
#include <chrono>
#include <ctime>

std::string getCETTimeForJSON()
{
    auto now = std::chrono::system_clock::now();
    auto now_time_t = std::chrono::system_clock::to_time_t(now);

    std::tm cet_tm = *std::gmtime(&now_time_t);

    std::stringstream ss;
    ss << std::put_time(&cet_tm, "%Y-%m-%dT%H:%M:%S+01:00");

    return ss.str();
}

int main()
{
    std::cout << "Czas CET (ISO8601): " << getCETTimeForJSON() << std::endl;
    return 0;
}